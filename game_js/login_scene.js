/*
    登录场景
 */

var LoginScene = Hilo.Class.create({
    Extends: Hilo.Container,
    constructor: function(properties){
        LoginScene.superclass.constructor.call(this, properties);
        this.init(properties);
    },
    init: function(properties){

        //text view
        var text = new Hilo.Text({
            font: "26px arial",
            text: "H5在线棋牌游戏 Demo",
            lineSpacing: 0,
            width: width,
            height: 100,
            x: 50,
            y: 50
        });

        var text1 = new Hilo.Text({
            font: "18px arial",
            text: "输入昵称登录游戏 : ",
            lineSpacing: 0,
            width: 250,
            height: 80,
            x: 50,
            y: 200
        });

        // 名称输入框
        var name = new Hilo.DOMElement({
            id:'name',
            element: Hilo.createElement('input', {
            }),
            width: 210,
            height: 36,
            x: 50,
            y: 230,
        });

        // 连接按钮
        var connBtn = new Hilo.Button({
            id: 'connBtn',
            image: 'resource/btn/btn_login.png',
            width: 88,
            height: 32,
            x: 50,
            y: 300
        });

        this.addChild(text, text1, name, connBtn);
    }
});
