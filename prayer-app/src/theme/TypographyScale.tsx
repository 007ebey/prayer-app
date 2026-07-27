import { typography } from "./tokens/typography";

const sample =
  "The quick brown fox jumps over the lazy dog.";

export default function TypographyScale() {
  return (
    <section className="space-y-8">

      <div>

        <h2 className="text-2xl font-bold">
          Typography Scale
        </h2>

        <p className="text-muted-foreground">
          Typography tokens used throughout the design system.
        </p>

      </div>

      <div className="overflow-hidden rounded-xl border">

        <table className="w-full border-collapse">

          <thead className="bg-muted">

            <tr>

              <th className="px-6 py-4 text-left">
                Variant
              </th>

              <th className="px-6 py-4 text-left">
                Preview
              </th>

              <th className="px-6 py-4 text-left">
                Size
              </th>

              <th className="px-6 py-4 text-left">
                Weight
              </th>

              <th className="px-6 py-4 text-left">
                Line Height
              </th>

              <th className="px-6 py-4 text-left">
                Font
              </th>

            </tr>

          </thead>

          <tbody>

            {Object.entries(
              typography.variants
            ).map(([name, value]) => (

              <tr
                key={name}
                className="border-t"
              >

                <td className="px-6 py-5 font-medium">
                  {name}
                </td>

                <td className="px-6 py-5">

                  <div
                    style={{
                      fontSize: value.fontSize,
                      fontWeight: value.fontWeight,
                      lineHeight: value.lineHeight,
                      fontFamily:
                        typography.fontFamily[
                          value.fontFamily
                        ],
                      textTransform:
                        value.textTransform,
                      letterSpacing:
                        value.letterSpacing,
                    }}
                  >
                    {sample}
                  </div>

                </td>

                <td className="px-6 py-5">
                  <code>{value.fontSize}</code>
                </td>

                <td className="px-6 py-5">
                  <code>{value.fontWeight}</code>
                </td>

                <td className="px-6 py-5">
                  <code>{value.lineHeight}</code>
                </td>

                <td className="px-6 py-5">
                  {value.fontFamily}
                </td>

              </tr>

            ))}

          </tbody>

        </table>

      </div>

    </section>
  );
}