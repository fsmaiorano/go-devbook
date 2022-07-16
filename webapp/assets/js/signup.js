$("#form-signup").on("submit", createUser);

function createUser(event) {
  debugger;
  event.preventDefault();

  if ($("#password").val() != $("#password-confirmation").val()) {
    // Swal.fire("Ops...", "Invalid password!", "error");
    return;
  }

  $.ajax({
    url: "/users",
    method: "POST",
    data: {
      name: $("#name").val(),
      email: $("#email").val(),
      nickname: $("#nickname").val(),
      password: $("#password").val(),
    },
  })
    .done(function () {
      debugger;
      //   Swal.fire("Sucess!", "User created with success!", "success").then(
      //     function () {
      //       $.ajax({
      //         url: "/login",
      //         method: "POST",
      //         data: {
      //           email: $("#email").val(),
      //           password: $("#password").val(),
      //         },
      //       })
      //         .done(function () {
      //           window.location = "/home";
      //         })
      //         .fail(function () {
      //           Swal.fire("Ops...", "Wrong user or password!", "error");
      //         });
      //     }
      //   );
    })
    .fail(function () {
      //   Swal.fire("Ops...", "User not created!", "error");
      debugger;
    });
}
