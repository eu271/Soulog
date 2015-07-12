var setTitulo = function(titulo){
    $('#tituloBlog').text(titulo);
}

var getBlog = function() {
    setTitulo("Cartas a un dios caido"); 
}


$(document).ready(function(){
    getBlog();

    blog.pedirPosts(10, vista.añadirPost);
});


