import { radius } from "./tokens/radius";

const values = [
  ["None", radius.none],
  ["XS", radius.xs],
  ["SM", radius.sm],
  ["MD", radius.md],
  ["LG", radius.lg],
  ["XL", radius.xl],
  ["2XL", radius["2xl"]],
  ["3XL", radius["3xl"]],
  ["Full", radius.full],
] as const;

export default function BorderRadius() {
  return (
    <section className="space-y-6">

      <div>
        <h2 className="text-xl font-semibold">
          Border Radius
        </h2>

        <p className="text-sm text-muted-foreground">
          Corner radius tokens used throughout the design system.
        </p>
      </div>

      <div className="grid gap-6 sm:grid-cols-2 lg:grid-cols-3">

        {values.map(([name, value]) => (

          <div
            key={name}
            className="space-y-3"
          >

            <div
              className="h-28 w-full border bg-background"
              style={{
                borderRadius: value,
              }}
            />

            <div className="space-y-1">

              <div className="font-medium">
                {name}
              </div>

              <code className="text-xs text-muted-foreground">
                {value}
              </code>

            </div>

          </div>

        ))}

      </div>

    </section>
  );
}