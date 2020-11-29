window.addEventListener("keydown", function (event) {
  if (event.defaultPrevented) {
    return; // Do nothing if the event was already processed
  }

  switch (event.key) {
    case "Down": 
    case "ArrowDown":
      prev();
      break;
    case "Up": 
    case "ArrowUp":
      next();
      break;
    case "Left": 
    case "ArrowLeft":
      prev();
      break;
    case "Right": // IE/Edge specific value
    case "ArrowRight":
      next();
      break;
    default:
      return; // Quit when this doesn't handle the key event.
  }

  // Cancel the default action to avoid it being handled twice
  event.preventDefault();
}, true);

function next() {
    if(window.location.hash) {
          let hash = window.location.hash.substring(1); //Puts hash in variable, and removes the # character
          window.location = "#" + (parseInt(hash) + 1);
          // hash found
    } else {
          window.location = "#" + 0;
    }
}

function prev() {
    if(window.location.hash) {
          let hash = window.location.hash.substring(1); //Puts hash in variable, and removes the # character
          window.location = "#" + (parseInt(hash) - 1);
          // hash found
    } else {
          window.location = "#" + 0;
    }
}
