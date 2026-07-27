// theme/Spacing.tsx

import { spacing } from "./tokens/spacing";

const usage: Record<string, string> = {
  "0": "No spacing",
  "1": "Icon padding",
  "2": "Compact gaps",
  "3": "Small component spacing",
  "4": "Default spacing",
  "5": "Comfortable spacing",
  "6": "Section spacing",
  "8": "Card padding",
  "10": "Large card spacing",
  "12": "Page sections",
  "16": "Hero spacing",
  "20": "Large layouts",
  "24": "Landing sections",
};

export default function Spacing() {
  return (
    <section className="space-y-8">

      <div>
        <h2 className="text-2xl font-bold">
          Spacing
        </h2>

        <p className="text-muted-foreground">
          Consistent spacing tokens used throughout the design system.
        </p>
      </div>

      <div className="space-y-6">

        {Object.entries(spacing).map(([token, value]) => (

          <div
            key={token}
            className="grid grid-cols-[80px_1fr_120px_220px] items-center gap-6"
          >

            <code className="font-semibold">
              {token}
            </code>

            <div className="flex items-center">

              <div
                className="h-6 rounded bg-primary"
                style={{
                  width: value,
                  minWidth: "1px",
                }}
              />

            </div>

            <code className="text-sm text-muted-foreground">
              {value}
            </code>

            <span className="text-sm text-muted-foreground">
              {usage[token]}
            </span>

          </div>

        ))}

      </div>

    </section>
  );
}