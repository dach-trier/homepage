export function registerHoverCapabilityController() {
    // lastTouchTime is used for ignoring emulated mousemove events
    let lastTouchTime = 0;

    function enableHover() {
        if (Date.now() - lastTouchTime < 500) return;
        document.body.classList.add("hasHover");
    }

    function disableHover() {
        document.body.classList.remove("hasHover");
    }

    function updateLastTouchTime() {
        lastTouchTime = Date.now();
    }

    document.addEventListener("touchstart", updateLastTouchTime, true);
    document.addEventListener("touchstart", disableHover, true);
    document.addEventListener("mousemove", enableHover, true);

    switch (document.readyState) {
        case "loading":
            document.addEventListener("DOMContentLoaded", enableHover);
            break;

        default:
            enableHover();
            break;
    }
}
