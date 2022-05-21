import { useParams } from "react-router-dom";
import NavBar from "../Components/Navbar";
import PaperForm from "../Components/PaperForm";

export default function WritePaperPage() {
  const { paperId } = useParams();
  return (
    <>
      <NavBar />
      <PaperForm paperId={paperId} />
    </>
  );
}
