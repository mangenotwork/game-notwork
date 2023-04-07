/*
炸金花游戏大厅
 */
var JinhuaScene = Hilo.Class.create({
    Extends: Hilo.Container,
    constructor: function(properties){
        JinhuaScene.superclass.constructor.call(this, properties);
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

        // 创建房间
        var newRoot = new Hilo.Button({
            id: 'newRoot',
            image: 'resource/btn/jinhua_root.png',
            width: 100,
            height: 64,
            x: 40,
            y: 320
        });

        // 房间列表 text view
        var rootListText = new Hilo.Text({
            font: "16px arial",
            text: "房间列表(点击进入房间)：",
            lineSpacing: 0,
            width: 100,
            height: 50,
            x: 40,
            y: 80
        });

        // 房间列表容器
        var rootList = new Hilo.Container({
            id:"rootList",
            x:40,
            y:100,
            background:"#aaa",
            width:width-40,
            height:200
        });

        this.addChild(text, returnHome, newRoot, rootListText, rootList);

    },

});