"use strict";

const Controller = require("egg").Controller;

class HomeController extends Controller {
  async index() {
    const { ctx, app } = this;
    const res = await app.redis.set("test", 1, "EX", 3600, "NX");
    console.log("res", res);
    ctx.body = "hi, egg";
  }
}

module.exports = HomeController;
