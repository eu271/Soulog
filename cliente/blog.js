var setTitulo = function(titulo){
    $('#tituloBlog').text(titulo);
}

var getBlog = function() {
    setTitulo("Cartas a un dios caido"); 
}


$(document).ready(function(){
    getBlog();

    blog.pedirPosts(10, vista.añadirPost);
    
    
    var templatePost = function(cantidad, funcmap) {
        $.ajax({
            method: "GET",
            url:"http://localhost:8080/templates/post.html.mustache"
            })
        .done(function(s){
            var post = {
                'Titulo': 'My interesting post',
                'Date': '0001-01-01T00:00:00Z',
                'Tags': [
                    {Tag:'interesting', TagLink:'interesting'},
                    {Tag:'my', TagLink:'my'}
                ],
                'Contenido': '[...] ## This is an interesting text about an interesting topic. [...]',
                
                
                
                'DateSimple': function(){ return 'returns a date in format yyyy-mm-dd'},
                'RenderDate': function(){ return 'returns a date in a human-readable format'},
                
                'RenderMarkdown': function(){ return 'returns "Contenido" render as html'}
            }
            
            var output = Mustache.render(s, post);
            
            $('#contenido').append(output);
        });
    };
    
    templatePost();
    
});


