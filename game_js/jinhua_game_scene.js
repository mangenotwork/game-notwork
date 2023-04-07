/*
    炸金花游戏房间
 */
var JinhuaGameScene = Hilo.Class.create({
    Extends: Hilo.Container,
    constructor: function (properties) {
        JinhuaGameScene.superclass.constructor.call(this, properties);
        this.init(properties);
    },

    init: function (properties) {

        var game_id = new Hilo.Container({
            id:"game_id",
        });

        // 房间头容器
        var game_title = new Hilo.Container({
            id:"game_title",
            x:20,
            y:20,
            background:"#fff",
            width:width,
            height:50
        });

        // 游戏退出按钮
        var game_out = new Hilo.Button({
            id: 'game_out',
            image: 'resource/btn/btn_out.png',
            width: 88,
            height: 36,
            x: 20,
            y: 60
        });

        // 游戏准备按钮
        var game_start = new Hilo.Button({
            id: 'game_start',
            image: 'resource/btn/btn_zb.png',
            width: 88,
            height: 36,
            x: 220,
            y: 60
        });

        // 玩家信息容器 - 主要显示玩家的准备信息
        var game_users = new Hilo.Container({
            id:"game_users",
            x:20,
            y:120,
            background:"#ddd",
            width:width-20,
            height:100
        });

        // 游戏的显示信息
        var game_show = new Hilo.Container({
            id:"game_show",
            x:20,
            y:230,
            background:"#ddd",
            width:width-20,
            height:120
        });

        // 玩家牌面
        var game_pai = new Hilo.Container({
            id:"game_pai",
            x:20,
            y:360,
            background:"#ddd",
            width:width-20,
            height:120
        });

        // 玩家操作区域
        var game_doing = new Hilo.Container({
            id:"game_doing",
            x:20,
            y:490,
            background:"#ddd",
            width:width-20,
            height:200
        });

        this.addChild(game_id, game_title, game_out, game_start, game_users, game_show, game_pai, game_doing);

    }
});