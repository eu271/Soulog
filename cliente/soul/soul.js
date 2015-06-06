var blog;
var secion = {
    "nombre": "Eugenio",
    "contraseña": "qwerty"
};
var post = new Array();



var enviarPost = function() {
    var post = {
        titulo : $('#tituloEditor').val(),
        contenido : $('#contenidoEditor').val(),
        secion : {
            nombre : secion.nombre,
            hash: secion.hash.toString()
        }
    };

    console.log(post);
    $.ajax({
        method: "POST",
        url:blog.url.sendPost,
        data: JSON.stringify(post) 
    })
        .done(function(respuesta){
            alert(respuesta); 
        });
}

var eliminarPost = function(clave_post) {
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
}


var pedirSecion = function(){
    
    var _peticion = {
        "Nombre": secion.nombre
    }

    $.ajax({
        method: "POST",
        url:blog.url.getSecion,
        data: JSON.stringify(_peticion) 
    })
        .done(function(s){
            s = JSON.parse(s);
            secion.secion = s.secion;
            secion.timestamp = s.timestamp;
            secion.hash = CryptoJS.SHA256(secion.nombre+secion.contraseña+secion.secion);
    })
}

var vista = {
    añadirPost : function(p){
        var stringHTML = '<div class="post" id="' + p.id + '">' +
            '<h2 class="titulo" id="titulo'+ p.id +'">' + p.titulo + '</h2>'+
            '<p class="contenido" id="contenido'+ p.id+ '">' +p.contenido + '</p></div>';

        $('#Posts').append(stringHTML);
    }
}

var ppp;

var pedirPosts = function() {
    $.ajax({
        method: "POST",
        url:blog.url.getPosts,
        data:'{"cantidad":1}'
    })
        .done(function(s){
            $.map(JSON.parse(s), function(v, i){
                 vista.añadirPost(v);
            });
        });
}

$(document).ready(function(){
    secion.contraseña = CryptoJS.SHA256(secion.contraseña).toString()

    pedirSecion();
});

