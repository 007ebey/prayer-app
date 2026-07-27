import { colors } from "./tokens/colors";
import { semantics } from "./tokens/semantic";

type ColorScale = Record<string, string>;

function ColorGroup({
  title,
  colors,
}: {
  title: string;
  colors: ColorScale;
}) {
  return (
    <section className="space-y-4">

      <h3 className="text-lg font-semibold">
        {title}
      </h3>

      <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-5">

        {Object.entries(colors).map(([name, value]) => (

          <div
            key={name}
            className="overflow-hidden rounded-xl border bg-background shadow-sm"
          >

            <div
              className="h-20"
              style={{
                backgroundColor: value,
              }}
            />

            <div className="space-y-1 p-3">

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

function PrimitivePalette({
  title,
  palette,
}: {
  title: string;
  palette: Record<number, string>;
}) {
  return (
    <section className="space-y-4">

      <h3 className="text-lg font-semibold">
        {title}
      </h3>

      <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-5">

        {Object.entries(palette).map(([shade, value]) => (

          <div
            key={shade}
            className="overflow-hidden rounded-xl border bg-background shadow-sm"
          >

            <div
              className="h-20"
              style={{
                backgroundColor: value,
              }}
            />

            <div className="space-y-1 p-3">

              <div className="font-medium">
                {shade}
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

export default function ColorPalette() {
  return (
    <div className="space-y-12">

      <div>

        <h2 className="text-2xl font-bold">
          Color Palette
        </h2>

        <p className="text-muted-foreground">
          Primitive colors and semantic design tokens used throughout the application.
        </p>

      </div>

      <PrimitivePalette
        title="Burgundy"
        palette={colors.primitive.burgundy}
      />

      <PrimitivePalette
        title="Gold"
        palette={colors.primitive.gold}
      />

      <PrimitivePalette
        title="Gray"
        palette={colors.primitive.gray}
      />

      <PrimitivePalette
        title="Red"
        palette={colors.primitive.red}
      />

      <PrimitivePalette
        title="Green"
        palette={colors.primitive.green}
      />

      <PrimitivePalette
        title="Blue"
        palette={colors.primitive.blue}
      />

      <PrimitivePalette
        title="Purple"
        palette={colors.primitive.purple}
      />

      <ColorGroup
        title="Brand"
        colors={semantics.brand}
      />

      <ColorGroup
        title="Background"
        colors={semantics.background}
      />

      <ColorGroup
        title="Surface"
        colors={semantics.surface}
      />

      <ColorGroup
        title="Text"
        colors={semantics.text}
      />

      <ColorGroup
        title="Border"
        colors={semantics.border}
      />

      <ColorGroup
        title="Status"
        colors={semantics.status}
      />

      <ColorGroup
        title="Session"
        colors={semantics.session}
      />

      <ColorGroup
        title="Prayer"
        colors={semantics.prayer}
      />

    </div>
  );
}