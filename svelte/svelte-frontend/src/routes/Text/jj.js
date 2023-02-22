export function call(){
   const sidenav2 = document.getElementById("sidenav-1");
const sidenavInstance2 = te.Sidenav.getInstance(sidenav2);

let innerWidth2 = null;

const setMode2 = (e) => {
  // Check necessary for Android devices
  if (window.innerWidth === innerWidth2) {
    return;
  }

  innerWidth2 = window.innerWidth;

  if (window.innerWidth < sidenavInstance2.getBreakpoint("xl")) {
    sidenavInstance2.changeMode("over");
    sidenavInstance2.hide();
  } else {
    sidenavInstance2.changeMode("side");
    sidenavInstance2.show();
  }
};

setMode2();

// Event listeners
window.addEventListener("resize", setMode2);
}