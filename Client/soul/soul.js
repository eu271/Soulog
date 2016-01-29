/* global $ */
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

var vistaAdmin = {
    "editor": false,
    "post": true,
    "ideas": false,


    mostrarEditor : function() {
        if( this.post ){
            this.post = false;
            this.editor = true;
            $('#posts').css('display', 'none');
            $('#editor').css('display', 'block');
        }
    },
    mostrarPosts: function() {
        if( this.editor ){
            this.post = true;
            this.editor = false;
            $('#posts').css('display', 'block');
            $('#editor').css('display', 'none');
        }
    },

    enviarPost:function(){
        var post = {
            titulo: $('#tituloEditor').val(),
            contenido: $('#contenidoEditor').val(),
        };
        blog.enviarPost(post);
    }

}

var reader = new commonmark.Parser();
var writer = new commonmark.HtmlRenderer();


$(document).ready(function(){
    secion.contraseña = CryptoJS.SHA256(secion.contraseña).toString();

    $("#editor-content").on('keypress', function(){
        $("#editor-preview").html(writer.render(reader.parse($("#editor-content").val())));

        var toTop = $("#editor-preview").prop('scrollHeight')*(0.01*((100*$("#editor-content").scrollTop())/$("#editor-content").prop('scrollHeight')));
        $("#editor-preview").scrollTop(toTop);
    })

    $("#editor-content").on('scroll', function(){
        var toTop = $("#editor-preview").prop('scrollHeight')*(0.01*((100*$("#editor-content").scrollTop())/$("#editor-content").prop('scrollHeight')));
        $("#editor-preview").scrollTop(toTop);
    })


    if(vistaAdmin.editor) {
        $("#Posts").css("display", "none");
    }else{
        $("#Editor").css("display", "none");
    }

    $('#menu-item-editor').click(function(){
        vistaAdmin.mostrarEditor();
    });
    $('#menu-item-posts').click(function(){
        vistaAdmin.mostrarPosts();
    });



});
