import { default as Web3 } from "web3";
import ecommerceStoreArtifacts from "../../build/contracts/EcommerceStore.json";

const ipfsAPI = require("ipfs-api");

const ipfs = ipfsAPI({ host: "localhost", port: "5001", protocol: "http" });

window.addEventListener("load", function () {
  if (window.ethereum) {
    // use MetaMask's provider
    App.web3 = new Web3(window.ethereum);
    window.ethereum.enable(); // get permission to access accounts
  } else {
    console.warn(
      "No web3 detected. Falling back to http://127.0.0.1:8545. You should remove this fallback when you deploy live"
    );
    // fallback - use your fallback strategy (local node / hosted node + in-dapp id mgmt / fail)
    App.web3 = new Web3(
      new Web3.providers.HttpProvider("http://127.0.0.1:8545")
    );
  }
  console.log("web3", App.web3);

  App.start();
});

window.App = {
  web3: null,
  account: null,
  meta: null,

  start: async function () {
    const { web3 } = this;
    try {
      // get contract instance
      const networkId = await web3.eth.net.getId();
      const deployedNetwork = ecommerceStoreArtifacts.networks[networkId];
      this.meta = new web3.eth.Contract(
        ecommerceStoreArtifacts.abi,
        deployedNetwork.address
      );

      // get accounts
      const accounts = await web3.eth.getAccounts();
      this.account = accounts[0];
      console.log("account", this.account);

      renderStore();
    } catch (err) {
      console.error("Could not connect to contract or chain.", err);
    }

    let reader;
    $("#product-image").change((event) => {
      const file = event.target.files[0];
      reader = new window.FileReader();
      reader.readAsArrayBuffer(file);
    });

    $("#add-item-to-store").submit(function (event) {
      const req = $("#add-item-to-store").serialize();
      console.log("reg:", req);

      let params = JSON.parse(
        '{"' +
          req.replace(/"/g, '\\"').replace(/&/g, '","').replace(/=/g, '":"') +
          '"}'
      );
      console.log("params", params);
      let decodedParams = {};
      Object.keys(params).forEach((key) => {
        decodedParams[key] = decodeURIComponent(decodeURI(params[key]));
      });
      saveProduct(reader, decodedParams);
      event.preventDefault();
    });

    if ($("#product-details").length > 0) {
      let productId = new URLSearchParams(window.location.search).get("id");
      renderProductDetails(productId);
    }

    $("#bidding").submit((event) => {
      $("#msg").hide();
      let amount = $("#bid-amount").val();
      let sendAmount = $("#bid-send-amount").val();
      let secretText = $("#secret-text").val();
      let productId = $("#product-id").val();
      let sealedBid = this.web3.utils.sha3(
        this.web3.utils.toWei(amount, "ether").toString() + secretText
      );
      console.log("sealedBid", productId, sealedBid);

      const { bid } = this.meta.methods;
      bid(parseInt(productId), sealedBid)
        .send({
          from: this.account,
          value: this.web3.utils.toWei(sendAmount, "ether"),
        })
        .then(() => {
          $("#msg").html("Your bid has been successfully submitted!");
          $("#msg").show();
        })
        .catch((err) => {
          console.error(err);
        });
      event.preventDefault();
    });

    $("#revealing").submit((event) => {
      $("#msg").hide();
      let amount = $("#actual-amount").val();
      let secretText = $("#reveal-secret-text").val();
      let productId = $("#product-id").val();

      const { revealBid } = this.meta.methods;
      revealBid(
        parseInt(productId),
        this.web3.utils.toWei(amount, "ether").toString(),
        secretText
      )
        .send({ from: this.account })
        .then((res) => {
          $("#msg").html("Your bid has been successfully revealed!");
          $("#msg").show();
        });
      event.preventDefault();
    });

    $("#finalize-auction").submit((event) => {
      $("#msg").hide();
      let productId = $("#product-id").val();

      const { finalizeAuction } = this.meta.methods;
      finalizeAuction(parseInt(productId))
        .send({ from: this.account })
        .then((res) => {
          $("#msg").html("The auction has been finalized and winner declared.");
          $("#msg").show();
          location.reload();
        })
        .catch((err) => {
          $("#msg").html(
            "The auction can not be finalized by the buyer or seller, only a third party aribiter can finalize it"
          );
          $("#msg").show();
        });
      event.preventDefault();
    });

    $("#release-funds").click(() => {
      let productId = new URLSearchParams(window.location.search).get("id");

      const { releaseAmountToSeller } = this.meta.methods;
      $("#msg")
        .html(
          "Your transaction has been submitted. Please wait for few seconds for the confirmation"
        )
        .show();
      releaseAmountToSeller(productId)
        .send({ from: this.account })
        .then(() => {
          location.reload();
        })
        .catch((err) => {
          console.log(err);
        });
    });

    $("#refund-funds").click(() => {
      let productId = new URLSearchParams(window.location.search).get("id");

      const { refundAmountToBuyer } = this.meta.methods;
      $("#msg")
        .html(
          "Your transaction has been submitted. Please wait for few seconds for the confirmation"
        )
        .show();
      refundAmountToBuyer(productId)
        .send({ from: this.account })
        .then((res) => {
          location.reload();
        })
        .catch((err) => {
          console.log();
        });
      alert("refund funds!");
    });
  },
};

function renderStore() {
  const { getProduct } = App.meta.methods;
  getProduct(1)
    .call()
    .then((p) => {
      $("#product-list").append(buildProduct(p));
    });
  getProduct(2)
    .call()
    .then((p) => {
      $("#product-list").append(buildProduct(p));
    });
}

function buildProduct(product) {
  let node = $("<div />");
  node.addClass("col-sm-3 text-center col-margin-bottom-1");
  node.append(
    "<a href='product.html?id=" +
      product[0] +
      "'><img src='http://localhost:8080/ipfs/" +
      product[3] +
      "' width='150px' /></a>"
  );
  node.append("<div>" + product[1] + "</div>");
  node.append("<div>" + product[2] + "</div>");
  node.append(
    "<div>" + new Date(product[5] * 1000).toLocaleString() + "</div>"
  );
  node.append(
    "<div>" + new Date(product[6] * 1000).toLocaleString() + "</div>"
  );
  node.append("<div> Ether " + displayPrice(product[7]) + "</div>");
  return node;
}

function saveProduct(reader, decodedParams) {
  let imageId, descId;
  saveImageOnIpfs(reader).then((id) => {
    imageId = id;
    saveTextBlobOnIpfs(decodedParams["product-description"]).then((id) => {
      descId = id;
      saveProductToBlockchain(decodedParams, imageId, descId);
    });
  });
}

function saveImageOnIpfs(reader) {
  return new Promise((resolve, reject) => {
    let buffer = Buffer.from(reader.result);
    ipfs
      .add(buffer)
      .then((res) => {
        console.log("image:", res);
        resolve(res[0].hash);
      })
      .catch((err) => {
        console.error(err);
        reject(err);
      });
  });
}

function saveTextBlobOnIpfs(blob) {
  return new Promise((resolve, reject) => {
    let buffer = Buffer.from(blob, "utf-8");
    ipfs
      .add(buffer)
      .then((res) => {
        console.log("res:", res);
        resolve(res[0].hash);
      })
      .catch((err) => {
        console.error(err);
        reject(err);
      });
  });
}

function saveProductToBlockchain(params, imageId, descId) {
  console.log("params in save product:", params);
  let auctionStartTime = Date.parse(params["product-auction-start"]) / 1000;
  let auctionEndTime =
    auctionStartTime + parseInt(params["product-auction-end"]) * 86400;

  const { addProductToStore } = App.meta.methods;
  addProductToStore(
    params["product-name"],
    params["product-category"],
    imageId,
    descId,
    auctionStartTime,
    auctionEndTime,
    App.web3.utils.toWei(params["product-price"], "ether"),
    parseInt(params["product-condition"])
  )
    .send({ from: App.account })
    .then(() => {
      $("#msg").show();
      $("#msg").html("Your product was successfully added to your store!");
    });
}

function renderProductDetails(productId) {
  const { getProduct, highestBidderInfo, escrowInfo } = App.meta.methods;
  console.log("id", productId);
  getProduct(productId)
    .call()
    .then((p) => {
      let desc = "";
      ipfs.cat(p[4]).then((file) => {
        desc = file.toString();
        $("#product-desc").append("<div>" + desc + "</div>");
      });
      $("#product-image").append(
        "<img src='http://localhost:8080/ipfs/" + p[3] + "' width='250px' />"
      );
      $("#product-name").html(p[1]);
      $("#product-price").html(displayPrice(p[7]));
      $("#product-id").val(p[0]);
      $("#product-auction-end").html(displayEndTime(p[6]));
      $("#bidding, #revealing, #finalize-auction, #escrow-info").hide();
      let currentTime = getCurrentTime();
      if (parseInt(p[8]) == 1) {
        $("#escrow-info").show();
        highestBidderInfo(productId)
          .call()
          .then((info) => {
            $("#product-status").html(
              "Auction has ended. Product sold to " +
                info[0] +
                " for " +
                displayPrice(info[2]) +
                "The money is in the escrow. Two of the three participants (Buyer, Seller and Arbiter) have to " +
                "either release the funds to seller or refund the money to the buyer"
            );
          });
        escrowInfo(productId)
          .call()
          .then((info) => {
            $("#seller").html("Seller: " + info[0]);
            $("#buyer").html("Buyer: " + info[1]);
            $("#arbiter").html("Arbiter: " + info[2]);
            if (info[3] == true) {
              $("#release-funds").hide();
              $("#refund-funds").hide();
              $("#release-count").html(
                "Amount from the escrow has been released"
              );
            } else {
              $("#release-count").html(
                info[4] +
                  " of 3 participants have agreed to release funds to seller"
              );
              $("#refund-count").html(
                info[5] + " of 3 participants have agreed to refund the buyer"
              );
            }
          });
      } else if (parseInt(p[8]) == 2)
        $("#product-status").html("Product not sold");
      else if (currentTime < p[6]) $("#bidding").show();
      else if (currentTime - 200 < p[6]) $("#revealing").show();
      else $("#finalize-auction").show();
    });
}

function displayPrice(amount) {
  return App.web3.utils.fromWei(amount, "ether") + " ETH";
}

function getCurrentTime() {
  return Math.round(new Date() / 1000);
}

function displayEndTime(timestamp) {
  let current_time = getCurrentTime();
  let remaining_time = timestamp - current_time;

  if (remaining_time <= 0) {
    return "Auction has ended";
  }

  let days = Math.trunc(remaining_time / 86400);
  remaining_time -= days * 86400;

  let hours = Math.trunc(remaining_time / 3600);
  remaining_time -= hours * 3600;

  let minutes = Math.trunc(remaining_time / 60);
  remaining_time -= minutes * 60;

  if (days > 0) {
    return (
      "Auction ends in " +
      days +
      " days" +
      hours +
      " hours" +
      minutes +
      " minutes" +
      remaining_time +
      " seconds"
    );
  } else if (hours > 0) {
    return (
      "Auction ends in " +
      hours +
      " hours" +
      minutes +
      " minutes" +
      remaining_time +
      " seconds"
    );
  } else if (minutes > 0) {
    return (
      "Auction ends in " + minutes + " minutes" + remaining_time + " seconds"
    );
  } else {
    return "Auction ends in " + remaining_time + " seconds";
  }
}
