var blog = {
    "titulo":"",
    "autor":""
};

var setTitulo = function(titulo){
    $('#tituloBlog').text(titulo);
}

var getBlog = function() {
    $.ajax({
        method:"POST",
        url: "/getSoul",
        data:{peticion:"getSoul"}
    })
        .done(function( data ) {
            blog = JSON.parse(data);
            setTitulo(blog.titulo);
        });
}


$(document).ready(function(){
    getBlog();
});


