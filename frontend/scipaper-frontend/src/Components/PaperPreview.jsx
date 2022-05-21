import Typography from "@mui/material/Typography";

export default function PaperPreview({ paper }) {
  return (
    <>
      <Typography variant="h2" align="center">
        {paper.title}
      </Typography>
      <Typography variant="h4" align="center">
        {paper.author} [{paper?.authorId}]
      </Typography>
      {paper.sections?.map((section, i) => (
        <div key={i}>
          <Typography variant="h5">{section.name}</Typography>
          <Typography>{section.content}</Typography>
        </div>
      ))}
    </>
  );
}
