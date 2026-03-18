import { registerHoverCapabilityController } from "@client/hover";
import { me } from "@client/inline";

if (window !== undefined) {
    // inline.ts exports
    (window as any).me = me;
}

registerHoverCapabilityController();
