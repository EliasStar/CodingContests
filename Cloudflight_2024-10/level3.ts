import { pack } from "npm:efficient-rect-packer";

const filenames = [ "level3/level3_example", "level3/level3_1", "level3/level3_2", "level3/level3_3", "level3/level3_4", "level3/level3_5" ]

for (const filename of filenames) {
    const input = Deno.readTextFileSync(filename + ".in")

    const rooms = []
    for (const line of input.split("\n").slice(1)) {
        if (line === "") continue
        const rawRoom = line.split(" ").map((x) => Number.parseInt(x))
        rooms.push({xLen: rawRoom[0], yLen: rawRoom[1], deskNum: rawRoom[2] })
    }

    let output = ""
    for (const room of rooms) {
        const rectangles = []
        for (let i = 1; i <= room.deskNum; i++) rectangles.push({ id: `${i}`, w: 3, h: 1 })

        const layout = await pack(rectangles, { w: room.xLen, h: room.yLen })

        const result = new Array(room.yLen).fill(0).map(() => new Array(room.xLen).fill("0"))
        for (const desk of layout.packed_rectangles) {
            for (let y = 0; y < desk.h; y++) {
                for (let x = 0; x < desk.w; x++) {
                    result[desk.y + y][desk.x + x] = desk.id
                }
            }
        }

        output += result.map(x => x.join(" ")).join("\n") + "\n\n"
    }

    Deno.writeTextFileSync(filename + ".out", output)
}
