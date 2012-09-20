/* http://www.html5rocks.com/en/tutorials/eventsource/basics/ */

if (!!window.EventSource) {
    var source = new EventSource('/blunder');
} else {
    // Result to xhr polling :(
}

source.addEventListener('message', function(e) {
    console.log('message', e);
    console.log(e.data);
}, false);

source.addEventListener('open', function(e) {
    // Connection was opened.
    console.log('opening');
}, false);

source.addEventListener('error', function(e) {
    if (e.readyState == EventSource.CLOSED) {
        // Connection was closed.
    }
    console.log('error', e);
}, false);
