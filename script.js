const fs = require("fs");
const csv = require("fast-csv");
const { format } = require("@fast-csv/format");

const data = [];

fs.createReadStream("./movies_metadata.csv")
  .pipe(csv.parse({ headers: true }))
  .on("error", (error) => console.error(error))
  .on("data", (row) => data.push(row))
  .on("end", async () => {
    await processData(data);
  });

async function processData(data) {
  for (let i = 0; i < data.length; i++) {
    let d = data[i];
    if (d["production_companies"] === "False") {
      d["production_companies"] = "[]";
    }
    if (d["genre"] === "False") {
      d["genre"] = "[]";
    }
    if (d["production_companies"] === "") {
      d["production_companies"] = "[]";
    }
    if (d["genre"] === "") {
      d["genre"] = "[]";
    }

    const productions = { productions: eval(d["production_companies"]) };
    const genres = { genres: eval(d["genres"]) };

    data[i]["genres"] = JSON.stringify(genres);
    data[i]["production_companies"] = JSON.stringify(productions);
  }

  const csvFile = fs.createWriteStream("movies_metadata.csv");
  const stream = format({ headers: true });
  stream.pipe(csvFile);

  stream.on("finish", function () {
    console.log("DONE!");
  });

  for (let d of data) {
    stream.write(d);
  }
  stream.end();
}
