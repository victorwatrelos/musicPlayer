$( document ).ready(function() {

    var url = 'localhost:8080';

    initBind();

    function initBind() {
        $('#launchSearch').on('click', function () {
            launchSearch();
        });

        $('#addToQueue').on('click', function() {
            addToQueue();
        });

        $('.emptyPlaylist').on('click', function() {
            sendJson('/clearQueue', {}, 'put', function() {});
        });

        $('.pauseMusic').on('click', function() {
            sendJson('/pause', {}, 'put', function() {});
        });

        $('.playMusic').on('click', function() {
            sendJson('/play', {}, 'put', function() {});
        });
    }

    function sendJson(url, data, type, callback) {
        $.ajax({
            url: url,
            type: type,
            data: JSON.stringify(data),
            contentType: "application/json; charset=utf-8",
            dataType: "json",
            success: callback
        });
    }

    function bindEvent() {
        $('.play').on('click', function () {
            var path = {'path': $(this).parent().data('path')};
            sendJson('/add', path, 'put', function() {
            });
        });

        $('.add').on('click', function () {
            var path = {'path': $(this).parent().data('path')};
            sendJson('/enqueue', path, 'put', function() {
            });
        });
    }

    function addToQueue() {
        var artist = $('#searchByArtist').val();
        sendJson('/addArtist', {'path': artist}, 'put', function() {});
    }

    function launchSearch() {
        var artist = $('#searchByArtist').val();
        $.get('/artist?q=' + artist, function(data, status){
            if (status != 'success') {
                console.log('Fail on /artist');
                console.log(status);
                return ;
            }
            var toAppend = $('#resultDisplay');
            toAppend.html('');
            var button = '<span class="glyphicon glyphicon-plus add" aria-hidden="true" ></span> ' + 
                        '<span class="glyphicon glyphicon-play play" aria-hidden="true" ></span>';
            if (data == null) {
                toAppend.append('<h2>No result</h2>');
                return ;
            }
            data.forEach(function(data) {
                var data = '<div class="list-group-item" data-path="' + data.Path + '">' +
                'Name: ' + data.Name + '<br/>' +
                'Album: ' + data.Album + ', Artiste: ' + data.Artist +
                button +
                '</div>';
                toAppend.append(data);
            });
            bindEvent();
        });
    }

});
