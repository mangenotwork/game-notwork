var Resource = Hilo.Class.create({
    Mixes: Hilo.EventMixin,

    queue: null,
    btn_returnhome : null,

    load: function(){
        var res = [
            {id:'btn_returnhome', src:'resource/btn/btn_returnhome.png'}
        ];

        this.queue = new Hilo.LoadQueue();
        this.queue.add(res);
        this.queue.on('complete', this.onComplete.bind(this));
        this.queue.start();
    },

    onComplete: function(e){
        this.btn_returnhome = this.queue.get('btn_returnhome').content;
        this.queue.off('complete');
        this.fire('complete');
    },

});