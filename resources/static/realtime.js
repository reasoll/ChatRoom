

function StartRealtime(roomid, timestamp) {

    StartSSE(roomid);
    StartForm();
}

function StartForm() {
    $('#chat-message').focus();
    $('#chat-form').ajaxForm(function() {
        $('#chat-message').val('');
        $('#chat-message').focus();
    });
}



function StartSSE(roomid) {
    if (!window.EventSource) {
        alert("EventSource is not enabled in this browser");
        return;
    }
    var source = new EventSource('/stream/'+roomid);
    source.addEventListener('message', newChatMessage, false);

}



function newChatMessage(e) {
    var data = jQuery.parseJSON(e.data);
    var nick = data.nick;
    var message = data.message;
    var style = rowStyle(nick);
    var html = "<tr class=\""+style+"\"><td>"+nick+"</td><td>"+message+"</td></tr>";
    $('#chat').append(html);

    $("#chat-scroll").scrollTop($("#chat-scroll")[0].scrollHeight);
}

function histogram(windowSize, timestamp) {
    var entries = new Array(windowSize);
    for(var i = 0; i < windowSize; i++) {
        entries[i] = {time: (timestamp-windowSize+i-1), y:0};
    }
    return entries;
}

var entityMap = {
    "&": "&amp;",
    "<": "&lt;",
    ">": "&gt;",
    '"': '&quot;',
    "'": '&#39;',
    "/": '&#x2F;'
};

function rowStyle(nick) {
    var classes = ['active', 'success', 'info', 'warning', 'danger'];
    var index = hashCode(nick)%5;
    return classes[index];
}

function hashCode(s){
  return Math.abs(s.split("").reduce(function(a,b){a=((a<<5)-a)+b.charCodeAt(0);return a&a},0));             
}

function escapeHtml(string) {
    return String(string).replace(/[&<>"'\/]/g, function (s) {
      return entityMap[s];
    });
}

window.StartRealtime = StartRealtime
