import { useEffect, useState } from "react";
import NavBar from "../Components/Navbar";
import PapersTable from "../Components/PapersTable";
import { myPapers } from "../Services/paperService";
import Typography from "@mui/material/Typography";
import { Box } from "@mui/system";

export default function HomePage() {
  const [papers, setPapers] = useState([]);

  useEffect(() => {
    myPapers().then((response) => {
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
        <Typography variant="h4">My papers</Typography>
        <PapersTable papers={papers} isMy={true} />
      </Box>
    </>
  );
}
