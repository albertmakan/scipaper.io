import { useEffect, useState } from "react";
import { toast } from "react-toastify";
import { createPaper, read, updatePaper } from "../Services/paperService";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import { Box } from "@mui/system";

export default function PaperForm({ paperId }) {
  const [paper, setPaper] = useState({
    title: "Title",
    sections: [
      {
        i: `${Math.random()}`,
        name: "",
        content: "",
      },
    ],
  });

  useEffect(() => {
    if (!paperId) return;
    read(paperId).then((p) => {
      console.log(p.title);
      p.sections?.forEach((s) => (s.i = `${Math.random()}`));
      setPaper(p);
    });
  }, [paperId]);

  const handleChange = (e, i) => {
    paper.sections.find((s) => s.i === i)[e.target.name] = e.target.value;
  };

  const handleAdd = (i) => {
    paper.sections.splice(paper.sections.findIndex((s) => s.i === i) + 1, 0, {
      i: `${Math.random()}`,
      name: "",
      content: "",
    });
    setPaper({
      ...paper,
      sections: paper.sections,
    });
  };

  const handleDelete = (i) => {
    setPaper({
      ...paper,
      sections: paper.sections.filter((s) => s.i !== i),
    });
  };

  const handleSave = () => {
    if (paper.id) {
      updatePaper(paper).then(() => toast.success("Paper updated"));
    } else {
      createPaper(paper).then((response) => {
        setPaper({ ...paper, id: response.id });
        toast.success("Paper is created");
      });
    }
  };

  const boxSx = {
    marginLeft: 15,
    marginRight: 15,
    marginTop: 5,
    marginBottom: 0,
  };

  return (
    <>
      <Box sx={boxSx}>
        <TextField
          label="Title"
          fullWidth
          value={paper.title}
          onChange={(e) => setPaper({ ...paper, title: e.target.value })}
          variant="filled"
        />
      </Box>
      {paper.sections.map((section) => (
        <Box key={section.i} sx={boxSx}>
          <TextField
            label="Section name"
            name="name"
            fullWidth
            defaultValue={section.name}
            onChange={(e) => handleChange(e, section.i)}
            variant="filled"
          />
          <TextField
            label="Content"
            name="content"
            fullWidth
            multiline
            rows={5}
            defaultValue={section.content}
            onChange={(e) => handleChange(e, section.i)}
            variant="filled"
          />
          <Button onClick={() => handleAdd(section.i)}>Insert section</Button>
          <Button onClick={() => handleDelete(section.i)}>Delete</Button>
        </Box>
      ))}
      <Box sx={boxSx}>
        <Button variant="contained" onClick={handleSave}>
          SAVE
        </Button>
      </Box>
    </>
  );
}
