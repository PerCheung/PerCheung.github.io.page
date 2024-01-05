点击返回[🔗我的博客文章目录](https://percheung.github.io/#/toc)
* 目录
{:toc}

# svg转png

> 写了一个spring boot项目，支持传入svg文件转出png图片，并且自定义转出png的宽和高。svg是矢量图，因此我们用svg可以转出任意大小的高清png大图。

## 主要代码

```java
import org.apache.batik.transcoder.Transcoder;
import org.apache.batik.transcoder.TranscoderException;
import org.apache.batik.transcoder.TranscoderInput;
import org.apache.batik.transcoder.TranscoderOutput;
import org.apache.batik.transcoder.image.PNGTranscoder;
import org.springframework.http.MediaType;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.multipart.MultipartFile;

import javax.servlet.http.HttpServletResponse;
import java.io.ByteArrayInputStream;
import java.io.IOException;
import java.io.InputStream;
import java.net.URLEncoder;
import java.nio.charset.StandardCharsets;
import java.util.Objects;

/**
 * @author Peter Cheung
 * 2023/2/16 15:21
 */
@RestController
public class SvgToPngController {
    /**
     * 上传svg
     * <p>
     * 然后
     * <p>
     * 下载png
     * <p>
     * consumes定义multipart/form-data
     */
    @GetMapping(path = "download", consumes = MediaType.MULTIPART_FORM_DATA_VALUE)
    public void download(float height, float width, MultipartFile file, HttpServletResponse response) throws IOException, TranscoderException {
        //MultipartFile转InputStream
        InputStream in = new ByteArrayInputStream(file.getBytes());
        Transcoder transcoder = new PNGTranscoder();
        //设置png图片的宽和长
        transcoder.addTranscodingHint(PNGTranscoder.KEY_WIDTH, width);
        transcoder.addTranscodingHint(PNGTranscoder.KEY_HEIGHT, height);
        try {
            TranscoderInput input = new TranscoderInput(in);
            //清空response
            response.reset();
            //强制下载不打开
            response.setContentType(MediaType.APPLICATION_OCTET_STREAM_VALUE);
            //设置编码为UTF_8
            response.setCharacterEncoding(StandardCharsets.UTF_8.name());
            //Content-Disposition的作用：告知浏览器以何种方式显示响应返回的文件，用浏览器打开还是以附件的形式下载到本地保存
            //attachment表示以附件方式下载 inline表示在线打开 "Content-Disposition:inline; filename=文件名.mp3"
            //filename表示文件的默认名称，因为网络传输只支持URL编码，因此需要将文件名URL编码后进行传输，前端收到后需要反编码才能获取到真正的名称
            response.setHeader("Content-Disposition", "attachment;filename=" + URLEncoder.encode((Objects.requireNonNull(file.getOriginalFilename()).split("\\."))[0], StandardCharsets.UTF_8.name()) + ".png");
            TranscoderOutput output = new TranscoderOutput(response.getOutputStream());
            transcoder.transcode(input, output);
        } finally {
            in.close();
        }
    }
}

```

## 所需依赖

```xml
<dependencies>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
        </dependency>
        <!--文件上传-->
        <dependency>
            <groupId>commons-fileupload</groupId>
            <artifactId>commons-fileupload</artifactId>
            <version>1.4</version>
        </dependency>
        <dependency>
            <groupId>commons-io</groupId>
            <artifactId>commons-io</artifactId>
            <version>2.11.0</version>
        </dependency>

        <dependency>
            <groupId>batik</groupId>
            <artifactId>batik-svggen</artifactId>
            <version>1.6</version>
        </dependency>
        <dependency>
            <groupId>batik</groupId>
            <artifactId>batik-awt-util</artifactId>
            <version>1.6</version>
        </dependency>
        <dependency>
            <groupId>batik</groupId>
            <artifactId>batik-bridge</artifactId>
            <version>1.6</version>
        </dependency>
        <dependency>
            <groupId>batik</groupId>
            <artifactId>batik-css</artifactId>
            <version>1.6</version>
        </dependency>
        <dependency>
            <groupId>batik</groupId>
            <artifactId>batik-dom</artifactId>
            <version>1.6</version>
        </dependency>
        <dependency>
            <groupId>batik</groupId>
            <artifactId>batik-gvt</artifactId>
            <version>1.6</version>
        </dependency>
        <dependency>
            <groupId>batik</groupId>
            <artifactId>batik-parser</artifactId>
            <version>1.6</version>
        </dependency>
        <dependency>
            <groupId>batik</groupId>
            <artifactId>batik-script</artifactId>
            <version>1.6</version>
        </dependency>
        <dependency>
            <groupId>batik</groupId>
            <artifactId>batik-svg-dom</artifactId>
            <version>1.6</version>
        </dependency>
        <dependency>
            <groupId>batik</groupId>
            <artifactId>batik-transcoder</artifactId>
            <version>1.6</version>
        </dependency>
        <dependency>
            <groupId>batik</groupId>
            <artifactId>batik-util</artifactId>
            <version>1.6</version>
        </dependency>
        <dependency>
            <groupId>batik</groupId>
            <artifactId>batik-xml</artifactId>
            <version>1.6</version>
        </dependency>
        <!-- 此处不能使用2.9.1版本，使用2.9.1生成png会失败 -->
        <dependency>
            <groupId>xerces</groupId>
            <artifactId>xercesImpl</artifactId>
            <version>2.5.0</version>
        </dependency>
        <dependency>
            <groupId>xml-apis</groupId>
            <artifactId>xmlParserAPIs</artifactId>
            <version>2.0.2</version>
        </dependency>

        <dependency>
            <groupId>org.axsl.org.w3c.dom.svg</groupId>
            <artifactId>svg-dom-java</artifactId>
            <version>1.1</version>
        </dependency>
        <dependency>
            <groupId>xml-apis</groupId>
            <artifactId>xml-apis</artifactId>
            <version>2.0.0</version>
        </dependency>

        <dependency>
            <groupId>org.w3c.css</groupId>
            <artifactId>sac</artifactId>
            <version>1.3</version>
        </dependency>


        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-test</artifactId>
            <scope>test</scope>
        </dependency>
</dependencies>
```

get请求，路径如同，body里使用form-data，三个参数，file是你要上传的svg文件，height是转出png图片的高，width是转出png图片的长，height和width是数字类型，支持整数和小数。

## 项目开源链接[🔗svg转png](https://github.com/PerCheung/svgtopng)

