console.log("woot crdt");

$(document).ready(function () {
  setInterval(function () {
    sync();
  }, 3000);
});

// Character Deleted
$("#text_box").on("keydown", function (event) {
  // If Backspace/Delete
  if (event.which === 8) {
    let position = $("#text_box").prop("selectionStart");

    console.log("Deleted: Position %s", position);

    let body = {
      position: position,
    };

    $.ajax({
      url: "/woot/delete",
      type: "POST",
      dataType: "json",
      data: JSON.stringify(body),
      contentType: "application/json; charset=utf-8",
    });
  }
});

// Character Added
$("#text_box").on("keypress", function (event) {
  let value = String.fromCharCode(event.which);
  let position = $("#text_box").prop("selectionStart") + 1;

  console.log("Added: Character %s Position %s", value, position);

  let body = {
    value: value,
    position: position,
  };

  $.ajax({
    url: "/woot/add",
    type: "POST",
    dataType: "json",
    data: JSON.stringify(body),
    contentType: "application/json; charset=utf-8",
  });
});

// Sync
function sync() {
  var data = $.ajax({
    url: "/woot/list",
    type: "GET",
    async: false,
  }).complete(function () {}).responseText;
  $("#text_box").val(data.replace(/\"/g, ""));
}
