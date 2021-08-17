import Web3 from "web3";
import votingArtifacts from "../../build/contracts/Voting.json";

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

  App.start();
});

window.App = {
  web3: null,
  account: null,
  meta: null,
  tokenPrice: null,
  candidates: null,
  address: null,

  start: async function () {
    const { web3 } = this;
    console.log("web3", web3);

    try {
      // get contract instance
      const networkId = await web3.eth.net.getId();
      const deployedNetwork = votingArtifacts.networks[networkId];
      this.address = deployedNetwork.address;
      console.log("contract address", this.address);
      this.meta = new web3.eth.Contract(votingArtifacts.abi, this.address);

      // get accounts
      const accounts = await web3.eth.getAccounts();
      this.account = accounts[0];
      console.log("account", accounts);

      populateCandidates();
    } catch (error) {
      console.error("Could not connect to contract or chain.", error);
    }
  },

  //购买投票
  buyTokens: function () {
    const { tokenSold, buy } = this.meta.methods;
    let tokensToBuy = $("#buy").val();
    let _value = tokensToBuy * this.tokenPrice;
    console.log("buy", buy, this.meta.methods);
    buy()
      .send({
        value: this.web3.utils.toWei(_value.toString(), "ether"),
        from: this.account,
      })
      .then(() => {
        tokenSold()
          .call()
          .then((amount) => {
            $("#tokens-sold").html(amount.toString());
          });
        this.web3.eth.getBalance(this.address, (err, balance) => {
          $("#contract-balance").html(
            this.web3.utils.fromWei(balance.toString(), "ether") + "ETH"
          );
        });
      });
  },

  //投票
  voteForCandidate: function () {
    const { voteForCandidate, totalVotesFor } = this.meta.methods;
    let candidateName = $("#candidate").val();
    let voteTokens = $("#vote-tokens").val();
    $("#candidate").val("");
    $("#vote-tokens").val("");
    voteForCandidate(this.web3.utils.asciiToHex(candidateName), voteTokens)
      .send({
        from: this.account,
      })
      .then(() => {
        totalVotesFor(this.web3.utils.asciiToHex(candidateName))
          .call()
          .then((count) => {
            $("#" + this.candidates[candidateName]).html(count.toString());
          });
      });
  },

  //投票人信息
  lookupVoterInfo: function () {
    const { voterDetails } = this.meta.methods;
    let _address = $("#voter-info").val();
    voterDetails(_address)
      .call()
      .then((res) => {
        $("#tokens-bought").html("Tokens Bought: " + res[0].toString());
        let candidateNames = Object.keys(this.candidates);
        $("#votes-cast").empty();
        $("#votes-cast").append("Votes cast per candidate: <br>");
        for (let i = 0; i < candidateNames.length; i++) {
          $("#votes-cast").append(
            candidateNames[i] + ": " + res[1][i].toString() + "<br>"
          );
        }
      });
  },
};

function populateCandidates() {
  const { allCandidate } = App.meta.methods;
  allCandidate()
    .call()
    .then((candidateArray) => {
      console.log("candidateArray", candidateArray);
      let candidates = {};
      for (let i = 0; i < candidateArray.length; i++) {
        candidates[App.web3.utils.toUtf8(candidateArray[i])] = "candidate-" + i;
      }
      App.candidates = candidates;
      setupCandidateRows();
      populateCandidateVotes();
      populateTokenData();
    });
}

function setupCandidateRows() {
  Object.keys(App.candidates).forEach((candidate) => {
    $("#candidate-rows").append(
      "<tr><td>" +
        candidate +
        "</td><td id='" +
        App.candidates[candidate] +
        "'></td></tr>"
    );
  });
}
function populateCandidateVotes() {
  const { totalVotesFor } = App.meta.methods;
  let candidateNames = Object.keys(App.candidates);
  console.log("candidateNames", candidateNames);
  for (let i = 0; i < candidateNames.length; i++) {
    totalVotesFor(App.web3.utils.asciiToHex(candidateNames[i]))
      .call()
      .then((count) => {
        $("#" + App.candidates[candidateNames[i]]).html(count.toString());
      });
  }
}
function populateTokenData() {
  const { totalTokens, tokenSold, tokenPrice } = App.meta.methods;
  totalTokens()
    .call()
    .then((amount) => {
      $("#tokens-total").html(amount.toString());
    });
  tokenSold()
    .call()
    .then((amount) => {
      $("#tokens-sold").html(amount.toString());
    });
  tokenPrice()
    .call()
    .then((price) => {
      console.log("price", price.toString());
      App.tokenPrice = App.web3.utils.fromWei(price.toString(), "ether");
      $("#token-cost").html(App.tokenPrice + "ETH");
    });
  App.web3.eth.getBalance(App.address).then((balance) => {
    console.log("balance", balance);
    $("#contract-balance").html(
      App.web3.utils.fromWei(balance.toString(), "ether") + "ETH"
    );
  });
}
