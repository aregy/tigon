
Console.WriteLine("Running document tasks");
var args1 = Environment.GetCommandLineArgs();
// foreach (var a in args) {
// 	Console.WriteLine(a);
// }
Console.WriteLine("---");
var fileName = args1[1];
var outPath = args1[2];
foreach (var a in args1) {
	Console.WriteLine(a);
}
var docProcessor = new DevExpress.Pdf.PdfDocumentProcessor();
docProcessor.LoadDocument(fileName);
for (var i = docProcessor.Document.Pages.Count; i > 1; i--){
	docProcessor.DeletePage(i);
}
docProcessor.SaveDocument(outPath);
