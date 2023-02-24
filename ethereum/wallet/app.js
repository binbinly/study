module.exports = (app) => {
  app.beforeStart(async () => {
    //初始化官方钱包
    await app.runSchedule("init_wallet");
  });
};
