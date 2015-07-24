/*
	Copyright (c) 2015 Eugenio Ochoa
	
	Permission is hereby granted, free of charge, to any person obtaining a copy
	of this software and associated documentation files (the "Software"), to deal
	in the Software without restriction, including without limitation the rights
	to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
	copies of the Software, and to permit persons to whom the Software is
	furnished to do so, subject to the following conditions:
	
	The above copyright notice and this permission notice shall be included in all
	copies or substantial portions of the Software.
	
	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
	SOFTWARE.
*/

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


