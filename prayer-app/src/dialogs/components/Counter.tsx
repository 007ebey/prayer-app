import { useState } from "react";

export default function Counter() {
    const [count, setCount] = useState<number>(0);

    return (
        <div className="bg-white rounded-xl shadow-lg p-8 w-80">
            <h2 className="text-2xl font-bold text-center">
                Counter
            </h2>

            <div className="text-center text-6xl my-8">
                {count}
            </div>

            <div className="flex justify-center gap-3">
                <button
                    className="bg-red-500 text-white px-4 py-2 rounded"
                    onClick={() => setCount(count - 1)}
                >
                    -
                </button>

                <button
                    className="bg-gray-600 text-white px-4 py-2 rounded"
                    onClick={() => setCount(0)}
                >
                    Reset
                </button>

                <button
                    className="bg-green-500 text-white px-4 py-2 rounded"
                    onClick={() => setCount(count + 1)}
                >
                    +
                </button>
            </div>
        </div>
    );
}