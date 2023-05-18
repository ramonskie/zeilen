// websocket.js

(function() {
    var socket = new WebSocket('ws://' + window.location.host + window.location.pathname);

    socket.onopen = function() {
      console.log('WebSocket connection established.');
    };

    socket.onmessage = function(event) {
      var message = JSON.parse(event.data);
      console.log('Received message:', message);

      // Update the HTML elements with the received data
      document.getElementById('playerID').textContent = message.playerID;
      document.getElementById('hand').textContent = message.hand;
      document.getElementById('points').textContent = message.points;
      document.getElementById('pot').textContent = message.pot;
    };

    socket.onclose = function(event) {
      console.log('WebSocket connection closed:', event);
    };
  })();
