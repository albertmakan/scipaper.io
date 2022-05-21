import { useEffect, useState } from "react";
import NavBar from "../Components/Navbar";
import PapersTable from "../Components/PapersTable";
import { getAll } from "../Services/libraryService";
import { Box } from "@mui/system";

export default function LibraryPage() {
  const [papers, setPapers] = useState([]);

  useEffect(() => {
    getAll().then((response) => {
      setPapers(response);
    });
  }, []);
  return (
    <>
      <NavBar />
      <Box
        sx={{
          marginLeft: 20,
          marginRight: 20,
          marginTop: 5,
          marginBottom: 0,
        }}
      >
        <PapersTable papers={papers} isMy={false} />
      </Box>
    </>
  );
}
