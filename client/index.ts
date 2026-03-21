import { registerHoverCapabilityController } from "@client/hover";
import { Highlights } from "@client/highlights";
import { me } from "@client/inline";

if (window !== undefined) {
    // inline.ts exports
    (window as any).me = me;

    // archive-carousel.ts exports
    (window as any).Highlights = Highlights;
}

registerHoverCapabilityController();
