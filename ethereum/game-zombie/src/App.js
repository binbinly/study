import React, { Component, Fragment } from "react";
import { Route, Switch, Redirect } from "react-router-dom";
import "./static/App.css";
import Nav from "./pages/Nav";
import Legion from "./pages/Legion";
import My from "./pages/My";
import Market from "./pages/Market";
import Gene from "./pages/Gene";
import Admin from "./pages/Admin";
import Detail from "./pages/Detail";
import Attack from "./pages/Attack";

class App extends Component {
  render() {
    return (
      <Fragment>
        <section className="zombies-hero no-webp block app-block-intro pt-5 pb-0">
          <div className="container">
            <div className="menu">
              <ul>
                <Nav />
              </ul>
            </div>
          </div>
        </section>
        <section className="zombie-container block bg-walls no-webp">
          <div className="container">
            <div className="area">
              <Switch>
                <Route path="/legion" component={Legion}></Route>
                <Route path="/my" component={My}></Route>
                <Route path="/market" component={Market}></Route>
                <Route path="/gene" component={Gene}></Route>
                <Route path="/admin" component={Admin}></Route>
                <Route path="/detail/:id" component={Detail}></Route>
                <Route path="/attack/:id" component={Attack}></Route>
                <Redirect to="/legion" />
              </Switch>
            </div>
          </div>
        </section>
      </Fragment>
    );
  }
}

export default App;
