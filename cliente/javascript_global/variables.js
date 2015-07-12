var secion = {
    "nombre": "Eugenio",
    "contraseña": "qwerty"
};

var blog = {
    url:{
        getPost : "/getPost",
        getTitulo: "/getTitulo",
        getSoul: "/getSoul",
        getPosts: "/getPosts",

        getSecion: "/getSecion",
        sendPost : "/sendPost",
        deletePost: "/deletePost"
    },

    pedirPosts:function(cantidad, funcmap) {
        $.ajax({
            method: "POST",
            url:blog.url.getPosts,
            data:'{"cantidad":'+cantidad+'}'
            })
        .done(function(s){
            $.map(JSON.parse(s), function(v, i){
                funcmap(v);
            });
        });
    },


// Llamadas que necesitan una peticion de secion para funcionar.
    pedirSecion:function(){
        var _peticion = {
            "Nombre": secion.nombre
        }

        return $.ajax({
                method: "POST",
                url:blog.url.getSecion,
                data: JSON.stringify(_peticion) 
            })
            .done(function(s){
                s = JSON.parse(s);
                secion.secion = s.secion;
                secion.timestamp = s.timestamp;
                secion.hash = CryptoJS.SHA256(secion.nombre+secion.contraseña+secion.secion);
            });
    },

    enviarPost:function(p) {
        blog.pedirSecion().done(function(){
            var post = {
                titulo : p.titulo,
                contenido : p.contenido,
                fechaPublicacion : p.fechaPublicacion,
                secion : {
                    nombre : secion.nombre,
                    hash: secion.hash.toString()
                }
            };

            $.ajax({
                method: "POST",
                url:blog.url.sendPost,
                data: JSON.stringify(post) 
            })
            .done(function(p){
                alert(p);
            });
        });
        
    },


    eliminarPost:function(clave_post) {
        blog.pedirSecion().done(function(){
            var paquete = {
                titulo : clave_post,
                secion :{
                    nombre : secion.nombre,
                    hash:secion.hash.toString()
                }
            }
            $.ajax({
                method: "POST",
                url:blog.url.deletePost,
                data:JSON.stringify(paquete)
            })
            .done(function(s){
                alert(s);
            });
        });
    }
};


    var vista = {
        añadirPost : function(p){
            var stringHTML = '<div class="post" id="' + p.id + '">' +
                '<h2 class="postTitulo" id="titulo'+ p.id +'">' + p.titulo + '</h2>'+
                '<p class="postContenido" id="contenido'+ p.id+ '">' +p.contenido + '</p>'+
                '<p class="postFecha" id="fecha'+ p.id+'">'+p.fechaPublicacion + '</p></div>';

            $('#contenido').append(stringHTML);
        }
    }
