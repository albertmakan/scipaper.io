import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import NavBar from "../Components/Navbar";
import PaperPreview from "../Components/PaperPreview";
import { read } from "../Services/paperService";

export default function PreviewPaperPage() {
  const { paperId } = useParams();
  const [paper, setPaper] = useState({ id: paperId });

  useEffect(() => {
    read(paperId).then((p) => {
      setPaper(p);
      if (!p) window.location.replace("/notfound");
    });
  }, [paperId]);

  return (
    <>
      <NavBar />
      <PaperPreview paper={paper} />
    </>
  );
}
