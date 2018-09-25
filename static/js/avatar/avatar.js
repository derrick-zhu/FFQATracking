function AvatarDrawCanvasWith(user_name, elementId) {
  var colours = ["#1abc9c", "#2ecc71", "#3498db", "#9b59b6", "#34495e", "#16a085", "#27ae60", "#2980b9", "#8e44ad",
    "#2c3e50", "#f1c40f", "#e67e22", "#e74c3c", "#95a5a6", "#f39c12", "#d35400", "#c0392b", "#bdc3c7", "#7f8c8d"
  ];

  var name = user_name;
  var nameSplit = name.split(" ");
  var initials;
  if (nameSplit.length < 2) {
    initials = name.charAt(0).toUpperCase() + name.charAt(1).toUpperCase()
  } else {
    initials = nameSplit[0].charAt(0).toUpperCase() + nameSplit[1].charAt(0).toUpperCase();
  }

  var charIndex = initials.charCodeAt(0) - 65,
    colourIndex = charIndex % 19;

  var allCanvas = document.getElementsByName(elementId);

  for (let eachCanvas of allCanvas) {

    var canvas = eachCanvas;
    var context = canvas.getContext("2d");

    var canvasWidth = canvas.width,
      canvasHeight = canvas.height,
      canvasCssWidth = canvasWidth,
      canvasCssHeight = canvasHeight;

    if (window.devicePixelRatio) {

      canvas.width = canvasWidth * window.devicePixelRatio;
      canvas.height = canvasHeight * window.devicePixelRatio;

      context.scale(window.devicePixelRatio, window.devicePixelRatio);
    }

    var font = (canvasWidth / 2).toString() + 'px Arial';

    context.fillStyle = colours[colourIndex];
    context.fillRect(0, 0, canvas.width, canvas.height);
    context.font = font;
    context.textAlign = "center";
    context.fillStyle = "#FFF";
    context.fillText(initials, canvasCssWidth / 2, canvasCssHeight / 1.5);
  }
}