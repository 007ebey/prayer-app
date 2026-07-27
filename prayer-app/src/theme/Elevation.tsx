// theme/Elevation.tsx

import { shadows } from "./tokens/shadows";
import { elevation } from "./tokens/elevation";

const items = [
  {
    name: "Flat",
    level: elevation.flat,
    shadow: "none",
    usage: "Backgrounds, dividers",
  },
  {
    name: "Raised",
    level: elevation.raised,
    shadow: shadows.sm,
    usage: "Cards, list items",
  },
  {
    name: "Floating",
    level: elevation.floating,
    shadow: shadows.md,
    usage: "Dropdowns, popovers",
  },
  {
    name: "Overlay",
    level: elevation.overlay,
    shadow: shadows.lg,
    usage: "Drawers, side panels",
  },
  {
    name: "Modal",
    level: elevation.modal,
    shadow: shadows.xl,
    usage: "Dialogs, modals",
  },
] as const;

export default function Elevation() {
  return (
    <section className="space-y-8">

      <div>
        <h2 className="text-2xl font-bold">
          Elevation
        </h2>

        <p className="text-muted-foreground">
          Elevation tokens define depth and visual hierarchy using shadows.
        </p>
      </div>

      <div className="grid gap-8 sm:grid-cols-2 xl:grid-cols-3">

        {items.map((item) => (

          <div
            key={item.name}
            className="space-y-4"
          >

            <div
              className="flex h-40 items-center justify-center rounded-xl border bg-white"
              style={{
                boxShadow: item.shadow,
              }}
            >
              <span className="font-medium">
                {item.name}
              </span>
            </div>

            <div className="space-y-1">

              <div className="flex items-center justify-between">

                <span className="font-medium">
                  {item.name}
                </span>

                <code className="text-xs text-muted-foreground">
                  {item.level}
                </code>

              </div>

              <div className="text-sm text-muted-foreground">
                {item.usage}
              </div>

            </div>

          </div>

        ))}

      </div>

    </section>
  );
}