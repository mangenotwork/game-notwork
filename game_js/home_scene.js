/*
    home场景
 */
var HomeScene = Hilo.Class.create({
    Extends: Hilo.Container,
    constructor: function(properties){
        HomeScene.superclass.constructor.call(this, properties);
        this.init(properties);
    },

    init: function(properties){

        //text view
        var text = new Hilo.Text({
            font: "26px arial",
            text: " > 主页",
            lineSpacing: 0,
            width: width,
            height: 50,
            x: 20,
            y: 20
        });

        var text1 = new Hilo.Text({
            font: "19px arial",
            text: "游戏列表 : ",
            width: width,
            height: 50,
            x: 20,
            y: 100
        });

        // 返回登录按钮
        var returnBtn = new Hilo.Button({
            id: 'returnBtn',
            image: 'resource/btn/btn_out.png',
            width: 88,
            height: 32,
            x: 20,
            y: height - 100
        });


        // 游戏1
        var gameBtn = new Hilo.Button({
            id: 'gameBtn',
            image: 'resource/btn/btn_game1.png',
            width: 90,
            height: 50,
            x: 40,
            y: 150
        });

        // pkp
        var pkpBtn = new Hilo.Button({
            id: 'pkpBtn',
            image: 'resource/btn/btn_pkp.png',
            width: 90,
            height: 50,
            x: 140,
            y: 150
        });

        //炸金花
        var jinhuaBtn = new Hilo.Button({
            id: 'jinhuaBtn',
            image: 'resource/btn/btn_jinhua.png',
            width: 90,
            height: 50,
            x: 240,
            y: 150
        });

        this.addChild(text, text1, returnBtn, gameBtn, pkpBtn, jinhuaBtn);
    }
});