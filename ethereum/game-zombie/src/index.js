import React from "react";
import ReactDOM from "react-dom";
import "./static/Index.css";
import App from "./App";
import { BrowserRouter } from "react-router-dom";
import zombie from "./utils/zombie";

React.$abi = new zombie();

ReactDOM.render(
  <BrowserRouter>
    <App />
  </BrowserRouter>,
  document.getElementById("root")
);
