import { z } from "zod";

const resSchema = z.object({
  message: z.string(),
});

type ResSchema = z.infer<typeof resSchema>;

export default async function Home() {
  const res: ResSchema = await fetch(
    (process.env.API_URL + "/hello") as string,
    {
      cache: "no-store",
    }
  )
    .then((res) => {
      if (!res.ok) {
        throw Error("Is API Running?");
      }
      return res.json();
    })
    .catch((e) => {
      console.log("parsing failed", e);
    });

  console.log(JSON.stringify(res));
  return (
    <div>
      <p>{JSON.stringify(res.message)}</p>
    </div>
  );
}
