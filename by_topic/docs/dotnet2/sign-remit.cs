using System.Drawing;
using System.Collections.Generic;
using DevExpress.Pdf;
using DevExpress.Drawing;
using System;
using System.IO;
using System.Security.Cryptography.X509Certificates;

using (PdfDocumentProcessor documentProcessor = new PdfDocumentProcessor()) {

    documentProcessor.LoadDocument(@"DevExpress Remittance Details.pdf");

    for (var i = documentProcessor.Document.Pages.Count; i > 1; i--){
	documentProcessor.DeletePage(i);
    }
    X509Certificate2 certificate = new X509Certificate2(@"pfx2.pfx", "", X509KeyStorageFlags.MachineKeySet);

    byte[] imageData = File.ReadAllBytes("sig.png");
    int pageNumber = 1;

    //int angleInDegrees = 45;
    double angleInRadians = 0;//angleInDegrees * (Math.PI / 180);
    PdfOrientedRectangle signatureBounds = new PdfOrientedRectangle(new PdfPoint(70, 240), 100, 130, angleInRadians);

//    PdfOrientedRectangle signatureBounds = new PdfOrientedRectangle(new PdfPoint(128, 250), 20, 40, angleInRadians);
    PdfSignature signature = new PdfSignature(certificate, imageData, pageNumber, signatureBounds);
    const float DrawingDpi = 72f;
    var content = DateTime.Now.ToString("MM/dd/yyyy");
    //using (DXSolidBrush textBrush = new DXSolidBrush(Color.FromArgb(255, Color.Black)))
    //                AddGraphics(documentProcessor, content, textBrush, DrawingDpi);
    documentProcessor.SaveDocument("..\\..\\RotatedDocumentWithGraphics.pdf");
    signature.Location = "USA";
    signature.ContactInfo = "aregy@devexpress.com";
    signature.Reason = "Approved";

    documentProcessor.SaveDocument(@"Remittance Details - Signed.pdf", new PdfSaveOptions() { Signature = signature });
}
static void AddGraphics(PdfDocumentProcessor processor, string text, DXSolidBrush textBrush, float DrawingDpi)
{
    IList<PdfPage> pages = processor.Document.Pages;
    for (int i = 0; i < pages.Count; i++) {
        PdfPage page = pages[i];
        using (PdfGraphics graphics = processor.CreateGraphics())
        {
            SizeF actualPageSize = PrepareGraphics(page, graphics, DrawingDpi, DrawingDpi);
            using (Font font = new Font("Segoe UI", 20, FontStyle.Regular))
            {
                SizeF textSize =
                    graphics.MeasureString(text, font, PdfStringFormat.GenericDefault);
                PointF topLeft = new PointF(0, 0);
                PointF bottomRight =
                    new PointF(actualPageSize.Width - textSize.Width-60, actualPageSize.Height - textSize.Height-215);
                //graphics.DrawString(text, font, textBrush, topLeft);
                graphics.DrawString(text, font, textBrush, bottomRight);
                graphics.AddToPageForeground(page, DrawingDpi, DrawingDpi);
            }
        }
    }
}

static SizeF PrepareGraphics(PdfPage page, PdfGraphics graphics, float dpiX, float dpiY) {
    PdfRectangle cropBox = page.CropBox;
    float cropBoxWidth = ConvertFromPdfUnits((float)cropBox.Width, dpiX);
    float cropBoxHeight = ConvertFromPdfUnits((float)cropBox.Height, dpiY);
    switch(page.Rotate) {
        case 90:
            graphics.RotateTransform(-90);
            graphics.TranslateTransform(-cropBoxHeight, 0);
            return new SizeF(cropBoxHeight, cropBoxWidth);
        case 180:
            graphics.RotateTransform(-180);
            graphics.TranslateTransform(-cropBoxWidth, -cropBoxHeight);
            return new SizeF(cropBoxWidth, cropBoxHeight);
        case 270:
            graphics.RotateTransform(-270);
            graphics.TranslateTransform(0, -cropBoxWidth);
            return new SizeF(cropBoxHeight, cropBoxWidth);
    }
    return new SizeF(cropBoxWidth, cropBoxHeight);
}

static float ConvertFromPdfUnits(float pdfValue, float targetDpi) {
    return pdfValue / 72f * targetDpi;
}
