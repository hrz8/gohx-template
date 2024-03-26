initNavbar();

function initNavbar() {
  const $navbarBurgers = document.querySelectorAll(".navbar-burger");
  function classToggler(event) {
    const target = event.target.closest("a") || event.target;

    const dataTarget = target.dataset.target;
    const $target = document.getElementById(dataTarget);

    target.classList.toggle("is-active");
    $target.classList.toggle("is-active");
  }

  $navbarBurgers.forEach((el) => {
    el.removeEventListener("click", classToggler);
    el.addEventListener("click", classToggler);
  });
}
