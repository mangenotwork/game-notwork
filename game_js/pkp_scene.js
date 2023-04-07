/*
扑克牌 pkp 场景
 */
var PkpScene = Hilo.Class.create({
    Extends: Hilo.Container,
    constructor: function(properties){
        Game1Scene.superclass.constructor.call(this, properties);
        this.init(properties);
    },

    init: function(properties){
        //text view
        var text = new Hilo.Text({
            font: "26px arial",
            text: "游戏: Game1",
            lineSpacing: 0,
            width: width,
            height: 50,
            x: 150,
            y: 20
        });

        // 返回首页
        var returnHome = new Hilo.Button({
            id: 'returnHome',
            image: resource.btn_returnhome,//'resource/btn/btn_returnhome.png',
            width: 88,
            height: 36,
            x: 20,
            y: 20
        });
        this.addChild(text, returnHome);

    },

});