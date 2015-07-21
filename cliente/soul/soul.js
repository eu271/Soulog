Date.prototype.addDays = function(days) 
{
    var dat = new Date(this.valueOf());
    dat.setDate(dat.getDate() + days);
    return dat;
}

var arrayDate = [[],[1,2,3,4,5,6,7],[1,2,3,4,5,6,7],[1,2,3,4,5,6,7],[1,2,3,4,5,6,7],[1,2,3,4,5,6,7]];

var nodelistNodes = [[],[1,2,3,4,5,6,7],[1,2,3,4,5,6,7],[1,2,3,4,5,6,7],[1,2,3,4,5,6,7],[1,2,3,4,5,6,7]];
var calendario = {
    'month': [
        'Ene','Feb','Mar','Abr','May','Jun','Jul','Ago','Sep','Oct','Nov','Dic' 
    ],
    'date': new Date(),
    'today': new Date(),
    'seleccionado': new Date(),

    limpiarCalendario: function(){
        var d = document.querySelectorAll(".seleccionado")[0];            
        if(d!=undefined){
            d.classList.remove("seleccionado");
        }

        var i=1;
        var j=0;
        var count = 0;
        var dateTem;
        for( var c=1; c<8; c++){
            document.querySelectorAll("#calendario-semana-1 #calendario-dia-"+c)[0].innerHTML = '';
            arrayDate[1][c] = dateTem;
        }
        var d = document.querySelectorAll(".today")[0];
        if( d!=undefined ) {
            d.classList.remove("today");
        }

    },
    
    setFunctions: function() {

        var hacerLlamada = function(i,j){
            return function(){
                var d = document.querySelectorAll(".seleccionado")[0];            
                if(d!=undefined){
                    d.classList.remove("seleccionado");

                }
                nodelistNodes[i][j].classList.add("seleccionado");
                calendario.seleccionado = arrayDate[i][j];
            }
        }
        
        for(var i=1;i<6;i++){
            for(var j=1;j<8;j++){
                nodelistNodes[i][j] = document.querySelectorAll("#calendario-semana-"+ i + " #calendario-dia-"+j)[0];
                nodelistNodes[i][j].addEventListener('click', hacerLlamada(i,j),true);
            }
        }
        
    },
    imprimirCalendario: function() {

        var i=1;
        var j=0;
        var count = 0;
        var dateTem;
        for( var c=1; c<8; c++){
            document.querySelectorAll("#calendario-semana-1 #calendario-dia-"+c)[0].innerHTML = '';
            arrayDate[1][c] = dateTem;
        }
        var d = document.querySelectorAll(".today")[0];
        if( d!=undefined ) {
            d.classList.remove("today");
        }
        var d = document.querySelectorAll(".seleccionado")[0];            
        if(d!=undefined){
            d.classList.remove("seleccionado");
        }

        calendario.date.setDate(1);
        
        document.querySelectorAll(".calendario-mes")[0].innerHTML =  calendario.month[calendario.date.getMonth()] + " " + calendario.date.getFullYear() ;


        
        for( var c=calendario.date.getDay()==0?7:calendario.date.getDay(); c<8; c++){
            j++;
            var casilla = document.querySelectorAll("#calendario-semana-"+i+" #calendario-dia-"+c)[0];
            dateTem = calendario.date.addDays(count);
            casilla.innerHTML = dateTem.getDate();

            dateTem.setHours(calendario.today.getHours(),
                    calendario.today.getMinutes(),
                    calendario.today.getSeconds(),
                    calendario.today.getMilliseconds());
            if ( dateTem.toString() == calendario.today.toString() ) {
                casilla.classList.add("today");
            }
            if ( dateTem.toString() == calendario.seleccionado.toString() ) {
                casilla.classList.add("seleccionado");
            }
            arrayDate[1][c] = dateTem;

            count++;
        }


        for( i=2; i<6; i++){
            for( j=1; j<8; j++){
                
                var casilla = document.querySelectorAll("#calendario-semana-"+i+" #calendario-dia-"+j)[0];
                dateTem = calendario.date.addDays(count);
                casilla.innerHTML = dateTem.getDate();

                dateTem.setHours(calendario.today.getHours(),
                    calendario.today.getMinutes(),
                    calendario.today.getSeconds(),
                    calendario.today.getMilliseconds());
                if ( dateTem.toString() == calendario.today.toString() ) {
                    casilla.classList.add("today");
                }
                if ( dateTem.toString() == calendario.seleccionado.toString() ) {
                    casilla.classList.add("seleccionado");
                }

                arrayDate[i][j] = dateTem;

                count++;
            } 
        }

    },

    setDate: function(d) {
        calendario.seleccion = new Date(d);
    },

    nextMonth: function(){
        calendario.date.setMonth((calendario.date.getMonth()+1));
        calendario.imprimirCalendario();
    },

    prevMonth: function(){
        calendario.date.setMonth((calendario.date.getMonth()-1));
        calendario.imprimirCalendario();
    }


};


var vistaAdmin = {
    "editor": true,
    "post": false,
    "ideas": false,


    mostrarEditor : function() {
        if( this.post ){
            this.post = false;
            this.editor = true;
            $('#Posts').css('display', 'none');
            $('#Editor').css('display', 'block');
        }
    },
    mostrarPosts: function() {
        if( this.editor ){
            this.post = true;
            this.editor = false;
            $('#Posts').css('display', 'block');
            $('#Editor').css('display', 'none');
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

    $("#contenidoEditor").on('keypress', function(){
        $("#editor-preview-post").html(writer.render(reader.parse($("#contenidoEditor").val())));
        
        var toTop = $("#editor-preview").prop('scrollHeight')
        *(0.01*((100*$("#contenidoEditor").scrollTop())
            /$("#contenidoEditor").prop('scrollHeight')));
        $("#editor-preview").scrollTop(toTop);
    });
    
    $("#contenidoEditor").on('scroll', function(){
        var toTop = $("#editor-preview").prop('scrollHeight')
        *(0.01*((100*$("#contenidoEditor").scrollTop())
            /$("#contenidoEditor").prop('scrollHeight')));
        $("#editor-preview").scrollTop(toTop);
    });
    
    calendario.imprimirCalendario(); 
    calendario.setFunctions();
    
    $("#Ideas").css("display", "none");

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

