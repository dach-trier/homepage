// prettier-ignore

export function Highlights(element: HTMLElement) {
    SlideShow(element);
    Parallax(element.querySelector(".highlight__image__container")!);
}

function SlideShow(element: HTMLElement) {
    const DURATION = 5000;

    let active_image: HTMLElement = element.querySelector(
        ".highlight__image[data-state='active']",
    )!;
    let active_label: HTMLElement = element.querySelector(
        ".highlight__label[data-state='active']",
    )!;
    let timeout;

    function getFirstChild(element: HTMLElement): HTMLElement {
        return element.children[0]! as HTMLElement;
    }

    // function getLastChild(element: HTMLElement): HTMLElement {
    //     const  children = element.children;
    //     const nchildren = element.children.length;

    //     return children[nchildren - 1]! as HTMLElement;
    // }

    // function getPreviousSibling(element: HTMLElement): HTMLElement {
    //     if (element.previousElementSibling)
    //         return element.previousElementSibling as HTMLElement;

    //     return getLastChild(element.parentElement! as HTMLElement);
    // }

    function getNextSibling(element: HTMLElement): HTMLElement {
        if (element.nextElementSibling)
            return element.nextElementSibling as HTMLElement;

        return getFirstChild(element.parentElement! as HTMLElement);
    }

    function selectNextSibling() {
        clearTimeout(timeout);

        active_image.dataset["state"] = "inactive";
        active_label.dataset["state"] = "inactive";

        const previous_image = getNextSibling(active_image);
        const previous_label = getNextSibling(active_label);

        active_image = previous_image;
        active_image.dataset["state"] = "active";
        active_label = previous_label;
        active_label.dataset["state"] = "active";

        setTimeout(selectNextSibling, DURATION);
    }

    setTimeout(selectNextSibling, DURATION);
}

function Parallax(element: HTMLElement) {
    window.addEventListener("scroll", function () {
        const progress: number = Math.min(
            1,
            window.scrollY / element.clientHeight,
        );

        element.style.transform = `scale(${1 + progress * 0.1})`;
        console.log(element.style.top);
    });
}
