/*
game1场景
 */
var Game1Scene = Hilo.Class.create({
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

        // 添加Game1 游戏素材
        var j = 0;
        var yHeight = 200;
        for (i=1;i<15;i++) {
            if (j > 4 ){
                j = 0;
                yHeight = yHeight + 100;
            }
            var tx = new Hilo.Button({
                id: 'e'+i,
                image: 'resource/tx/e'+i+'.ico',
                width: 64,
                height: 64,
                x: 20+j*70,
                y: yHeight
            });
            // 每个素材的操作事件
            tx.on(Hilo.event.POINTER_START, function(e){
                game_Game1(this.id);
            }).on(Hilo.event.POINTER_END, function(e){
            });
            this.addChild(tx);
            j++;
        }
    },

});