点击返回[🔗我的博客文章目录](https://percheung.github.io/#/toc)
* 目录
{:toc}
<div onclick="window.scrollTo({top:0,behavior:'smooth'});" style="background-color:white;position:fixed;bottom:20px;right:40px;padding:10px 10px 5px 10px;cursor:pointer;z-index:10;border-radius:13%;box-shadow:0.5px 3px 7px rgba(0,0,0,0.3);"><img src="https://percheung.github.io/blogImg/backTop.png" alt="TOP" style="background-color:white;width:30px;"></div>

# 前后端分离使用<img src="https://percheung.github.io/blogImg/sa-token.png" width="70px" alt="sa-token" />Sa-Token（超越官方文档）

## 前言

Sa-Token的官方文档链接[🔗https://sa-token.cc/doc.html](https://sa-token.cc/doc.html)在此。

**事先声明**，起一个这样的标题并不是我狂妄自大，而且Sa-Token的官方文档是我见过的少有的写的很好的官方文档（很多知名项目的官方文档可以说一言难尽）。但是，关于前后端分离的这部分，我感到Sa-Token官方文档实在有些过于简略。而我在学习Sa-Token的时候也遇到一个困境，就是关于前后端分离的文章在百度谷歌上基本找不来，大部分所谓原创都是对官方文档的复制粘贴，照搬。因此，在我掌握前后端分离对Sa-Token的使用后，我决定写这样一篇文章，方便后来者对Sa-Token的前后端分离的使用有一个参考。

在**前后端分离**的权限验证，登录校验这条路上。我自学习编程以来，先后经历过传统token，jwt，单点登录，shiro，Spring Security等。这些技术，有的需要你书写大篇幅的拦截器，过滤器，资源类控制逻辑。可谓苦不堪言。而且官方文档一塌糊涂，使用步骤极度繁琐。最终，我找到了一个目前我认为最完美的技术：Sa-Token。当今国内企业，大部分项目都是前后端分离的，Sa-Token美中不足的地方在于，它的官方文档对前后端不分离的部分讲述的很清楚，对我们真正要经常用到前后端分离的使用讲述的却过于简略。这篇文档，我将从头到尾讲解如何使用Sa-Token做前后端分离的**Spring Boot**项目。

请注意，大部分代码需要你结合实际，甚至需要你稍加改动才能用。我会默认你已经掌握了jwt，cookie，session等知识。我将在每一段代码内写上注释，而且讲解我的写作逻辑，但这不意味着你可以不思考就能掌握这项技术。我已经将我的代码开源放在github，你也可以拉取GitHub的代码在本地运行，更方便你理解这项技术。**GitHub**的链接[🔗https://github.com/PerCheung/learnsatoken](https://github.com/PerCheung/learnsatoken)在此。

## 步骤

### 1.创建一个spring boot项目，引入如下两个依赖。

```xml
        <!-- Sa-Token 权限认证，在线文档：https://sa-token.cc -->
        <dependency>
            <groupId>cn.dev33</groupId>
            <artifactId>sa-token-spring-boot-starter</artifactId>
            <version>1.35.0.RC</version>
        </dependency>
        <!-- Sa-Token 整合 Redis （使用 jackson 序列化方式） -->
        <dependency>
            <groupId>cn.dev33</groupId>
            <artifactId>sa-token-redis-jackson</artifactId>
            <version>1.35.0.RC</version>
        </dependency>
```
讲解：第一个是sa-token的自动装配，无需多言。第二个，则会将我们的token自动管理在redis中，而不是在cookie或者session，这样就达到了最完美的前后端分离状态，无论你是重启前端还是后端，销毁cookie还是session，都不影响项目的token。如果你对单点登录，分布式登录有了解，你就会深深明白这么做的好处。在跨域的单点登录里，cookie是失效的，前后端分离，分布式的项目里，一个请求根本不知道到来的请求在跟哪个session打交道。**我的代码只求最彻底，最激进，最完善的前后端分离**。
### 2.配置你的application.yml

> 这里需要注意，在spring boot data redis的2版本以后，已经弃用 jedis 改用 lettuce，而官网仍在配置jedis连接池，这里要用lettuce连接池。

```yaml
############## Sa-Token 配置 (文档: https://sa-token.cc) ##############
sa-token:
  # token 名称（同时也是 cookie 名称）
  token-name: token
  # token 有效期（单位：秒） 默认30天，-1 代表永久有效
  timeout: -1
  # token 最低活跃频率（单位：秒），如果 token 超过此时间没有访问系统就会被冻结，默认-1 代表不限制，永不冻结
  active-timeout: -1
  # 是否允许同一账号多地同时登录 （为 true 时允许一起登录, 为 false 时新登录挤掉旧登录）
  is-concurrent: false
  # token 风格（默认可取值：uuid、simple-uuid、random-32、random-64、random-128、tik）
  token-style: simple-uuid
  # 是否输出操作日志
  is-log: true
  # 是否尝试从 cookie 里读取 Token，此值为 false 后，StpUtil.login(id) 登录时也不会再往前端注入Cookie
  isReadCookie: false
spring:
  # Redis配置
  redis:
    host: localhost
    port: 6379
    # 根据自己设置的密码决定
    password: 你的redis密码
    # 操作0号数据库，默认有16个数据库
    database: 0
    lettuce:
      pool:
        # 最大连接数
        max-active: 500
        # 连接池最大阻塞等待时间
        max-wait: 1000ms
        # 连接池中的最大空闲连接
        max-idle: 100
        # 连接池中的最小空闲连接
        min-idle: 0
```

讲解：如果你用过jwt，那么**token-name: token**就可以达到类似于jwt的效果，在你的header加上**token:token值**，就可以达到登录的校验的效果，像jwt一样，完全抛弃了cookie和session，而jwt过于冗余，**token-style: simple-uuid**将保证token的精简，结合token的唯一性配合redis的使用，jwt携带信息的优点也将不复存在，因为redis可以存储一切。而redis配置你不必担心它影响spring boot自带的redis自动装配依赖，结合实际你会发现sa-token就是封装的redis自动装配依赖。当你创造一个token，他也会自动写入redis，不需要你多写任何代码。所有的token，都将会自动管理。

### 3.注解鉴权

加sa-token拦截器。

```java
import cn.dev33.satoken.interceptor.SaInterceptor;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.servlet.config.annotation.InterceptorRegistry;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer;

/**
 * @author Peter Cheung
 * 2023/7/19 13:52
 */
@Configuration
public class SaTokenConfigure implements WebMvcConfigurer {
    // 注册 Sa-Token 拦截器，打开注解式鉴权功能
    @Override
    public void addInterceptors(InterceptorRegistry registry) {
        // 注册 Sa-Token 拦截器，打开注解式鉴权功能
        registry.addInterceptor(new SaInterceptor()).addPathPatterns("/api/**");
    }
}
```
讲解：/api最好写上，然后把你的接口都带上/api，因为/**什么都扫描，效率低下。**如果你只做登录校验，这么点代码就已经足够。**

如果你还想做权限，角色的校验，加上下面的代码。

```java
import cn.dev33.satoken.stp.StpInterface;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;

/**
 * @author Peter Cheung
 * 2023/7/19 16:55
 */
@Service
public class StpInterfaceImpl implements StpInterface {

    /**
     * 返回一个账号所拥有的权限码集合
     */
    @Override
    public List<String> getPermissionList(Object loginId, String loginType) {
        // 本 list 仅做模拟，实际项目中要根据具体业务逻辑来查询权限
        List<String> list = new ArrayList<String>();
        list.add("101");
        list.add("user.add");
        list.add("user.update");
        list.add("user.get");
        // list.add("user.delete");
        list.add("art.*");
        return list;
    }

    /**
     * 返回一个账号所拥有的角色标识集合 (权限与角色可分开校验)
     */
    @Override
    public List<String> getRoleList(Object loginId, String loginType) {
        // 本 list 仅做模拟，实际项目中要根据具体业务逻辑来查询角色
        List<String> list = new ArrayList<String>();
        if ("root".equals(loginId)) {
            list.add("admin");
            list.add("super-admin");
        }
        return list;
    }
}

```
讲解：具体用法我不再赘述，如果你学过shiro，那你理解这段代码轻而易举，无非是把一个角色数组或者权限数据，跟账号做了一个绑定。如果你真的想理解掌握，然后超越，去改造这段代码，需要你不得不去学习一下shiro的五表思想去分配角色和权限资源。才能让你真正在实际工作里拿捏对权限角色的分配管理。

### 4.加入RedisConfig

这和sa-token无关，只是为了让你的redis能支持中文存储。

```java
import org.springframework.cache.annotation.CachingConfigurerSupport;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.data.redis.connection.RedisConnectionFactory;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.data.redis.serializer.GenericJackson2JsonRedisSerializer;
import org.springframework.data.redis.serializer.StringRedisSerializer;

/**
 * Redis配置类，更换默认序列化器
 *
 * @author Peter Cheung
 * @since 2023-07-28 13:51:55
 */
@Configuration
public class RedisConfig extends CachingConfigurerSupport {
    /**
     * 创建 RedisTemplate Bean，用于操作 Redis 数据库。
     *
     * @param connectionFactory Redis 连接工厂
     * @return RedisTemplate 实例
     */
    @Bean
    public RedisTemplate<String, Object> redisTemplate(RedisConnectionFactory connectionFactory) {
        RedisTemplate<String, Object> redisTemplate = new RedisTemplate<>();
        // 设置 RedisTemplate 的连接工厂
        redisTemplate.setConnectionFactory(connectionFactory);
        // 创建 StringRedisSerializer 实例，用于序列化和反序列化 Redis 的键和哈希键
        StringRedisSerializer stringRedisSerializer = new StringRedisSerializer();
        // 创建 GenericJackson2JsonRedisSerializer 实例，用于序列化和反序列化 Redis 的值和哈希值
        GenericJackson2JsonRedisSerializer genericJackson2JsonRedisSerializer = new GenericJackson2JsonRedisSerializer();
        // 设置默认序列化器为 StringRedisSerializer
        redisTemplate.setDefaultSerializer(stringRedisSerializer);
        // 设置 RedisTemplate 的键序列化器为 StringRedisSerializer
        redisTemplate.setKeySerializer(stringRedisSerializer);
        // 设置 RedisTemplate 的哈希键序列化器为 StringRedisSerializer
        redisTemplate.setHashKeySerializer(stringRedisSerializer);
        // 设置 RedisTemplate 的值序列化器为 GenericJackson2JsonRedisSerializer
        redisTemplate.setValueSerializer(genericJackson2JsonRedisSerializer);
        // 设置 RedisTemplate 的哈希值序列化器为 GenericJackson2JsonRedisSerializer
        redisTemplate.setHashValueSerializer(genericJackson2JsonRedisSerializer);
        return redisTemplate;
    }
}

```

### 5.异常统一处理

如果你的Java知识不足以理解这些代码，这些也可以不要，下面的代码只是为了让你错误处理更加优雅。

```java
import javax.validation.ValidationException;
import java.io.ByteArrayOutputStream;
import java.io.PrintStream;

import static com.learn.learnsatoken.config.constant.Constant.PACKAGE_NAME;

/**
 * 全局异常统一处理
 *
 * @author Peter Cheung
 * @since 2023-07-19 11:37:24
 */
@Slf4j
@ControllerAdvice
@ResponseBody
public class AllExceptionHandle {
    /**
     * 登录权限校验
     */
    @ExceptionHandler({NotLoginException.class, NotRoleException.class})
    public ResponseEntity<R> unauthorized(Exception e) {
        return R.deal(R.unauthorized().data(e(e)));
    }

    /**
     * 校验传参
     */
    @ExceptionHandler(ValidationException.class)
    public ResponseEntity<R> handleBadRequest(Exception e) {
        return R.deal(R.badRequest().data(e(e)));
    }

    /**
     * 全局异常处理
     */
    @ExceptionHandler(Exception.class)
    public ResponseEntity<R> exception(Exception e) {
        return R.deal(R.exp().data(e(e)));
    }

    /**
     * 异常信息处理主体方法
     *
     * @param e 异常对象
     * @return 异常解析信息
     */
    private String e(Exception e) {
        ByteArrayOutputStream printStackTrace = new ByteArrayOutputStream();
        e.printStackTrace(new PrintStream(printStackTrace));
        log.error(String.valueOf(printStackTrace));
        //错误信息
        StringBuilder errorMessage = new StringBuilder();
        errorMessage.append(e);
        if (StringUtils.isBlank(e.getMessage())) {
            //处理
            log.error(String.valueOf(errorMessage));
            return String.valueOf(errorMessage);
        }
        StackTraceElement[] stackTrace = e.getStackTrace();
        for (StackTraceElement stackTraceElement : stackTrace) {
            String className = stackTraceElement.getClassName();
            if (className.startsWith(PACKAGE_NAME)) {
                String errorName = ";" + stackTraceElement.getClassName();
                errorMessage.append(errorName);
                String errorLineNumber = ":" + stackTraceElement.getLineNumber();
                errorMessage.append(errorLineNumber);
                //处理
                log.error(String.valueOf(errorMessage));
                return String.valueOf(errorMessage);
            }
        }
        //处理
        log.error(String.valueOf(errorMessage));
        return String.valueOf(errorMessage);
    }
}
```

你只需关注的代码是这一段。

```java
    /**
     * 登录权限校验
     */
    @ExceptionHandler({NotLoginException.class, NotRoleException.class})
    public ResponseEntity<R> unauthorized(Exception e) {
        return R.deal(R.unauthorized().data(e(e)));
    }
```
讲解：登录，权限的不足，都可以理解成401错误，这是http的知识，这段代码将会把sa-token的报错都以401错误返回给前端，让前端一看便知：**哦，原来是用户的问题，不是后端错了。**

### 6.业务代码controller层

```java
import cn.dev33.satoken.annotation.SaCheckBasic;
import cn.dev33.satoken.annotation.SaCheckLogin;
import cn.dev33.satoken.annotation.SaCheckRole;
import cn.dev33.satoken.stp.StpUtil;
import cn.dev33.satoken.util.SaResult;
import com.learn.learnsatoken.mvc.domain.User;
import com.learn.learnsatoken.mvc.service.UserService;
import com.learn.learnsatoken.util.MD5Util;
import com.learn.learnsatoken.util.R;
import com.learn.learnsatoken.util.UUIDUtil;
import io.swagger.annotations.Api;
import io.swagger.annotations.ApiOperation;
import io.swagger.annotations.ApiParam;
import lombok.extern.slf4j.Slf4j;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import javax.annotation.Resource;
import java.util.concurrent.TimeUnit;

/**
 * 用户表(User)表控制层
 *
 * @author Peter Cheung
 * @since 2023-07-19 11:37:23
 */
@CrossOrigin(origins = "*", methods = {RequestMethod.GET, RequestMethod.POST, RequestMethod.PUT, RequestMethod.DELETE, RequestMethod.HEAD})
@Slf4j
@RestController
@RequestMapping("api/user")
@Api(tags = "用户表(User)表控制层")
public class UserController {
    /**
     * 服务对象
     */
    @Resource
    private UserService userService;

    @Resource
    private RedisTemplate<String, String> redisTemplate;

    /**
     * 注册
     */
    @GetMapping("register")
    public ResponseEntity<R> register(User user) {
        String password = user.getPassword();
        String salt = UUIDUtil.toUUID();
        user.setSalt(salt);
        user.setPassword(MD5Util.toMD5(password, salt));
        return R.deal(this.userService.insert(user));
    }

    @GetMapping("login")
    public ResponseEntity<R> doLogin(User user) {
        String username = user.getUsername();
        User data = (User) userService.queryById(username).getData();
        if (data == null) {
            return R.deal(R.unauthorized().data("no user"));
        }
        if (data.getPassword().equals(MD5Util.toMD5(user.getPassword(), data.getSalt()))) {
            StpUtil.login(username);
            return R.deal(R.ok().data(StpUtil.getTokenInfo().getTokenValue()));
        }
        return R.deal(R.unauthorized().data("wrong password"));
    }

    @GetMapping("logout")
    public SaResult logout(String name, String token) {
        redisTemplate.opsForValue().set("最近五分钟访问的ip", "我是？123456", 5, TimeUnit.MINUTES);
        //StpUtil.logout(name);
        StpUtil.logoutByTokenValue(token);
        return SaResult.ok();
    }

    /**
     * 全查询
     *
     * @param user 筛选条件
     * @return 查询结果
     */
    @SaCheckLogin
    @ApiOperation("全查询")
    @GetMapping
    public ResponseEntity<R> queryAll(@ApiParam(value = "user 筛选条件") User user) {
        return R.deal(this.userService.queryAll(user));
    }

    @SaCheckBasic(account = "admin:admin")
    @GetMapping("basic")
    public ResponseEntity<R> basic(String token) {
        Object username = StpUtil.getLoginIdByToken(token);
        return R.deal(R.ok().data(username));
    }

    /**
     * 通过主键查询单条数据
     *
     * @param id 主键
     * @return 单条数据
     */
    @SaCheckRole("admin")
    @ApiOperation("通过主键查询单条数据")
    @GetMapping("{id}")
    public ResponseEntity<R> queryById(@ApiParam(value = "id 主键") @PathVariable("id") String id) {
        return R.deal(this.userService.queryById(id));
    }
}
```

讲解：你要注意到，这里面有几个注解`@SaCheckLogin`，`@SaCheckRole("admin")`，第一个就是登录校验，第二个就是admin角色校验。你观察我的login部分。`StpUtil.login(username);`这句话，就能保证token以username为key，创造token写入redis并且管理起来，完全无需太多代码。而`StpUtil.getTokenInfo().getTokenValue()`将把token拿出来返回给前端，前端要记得把token保存起来，最好是放前端的**localStorage**（后端不懂这是什么，前端很懂）里，只要前端记得后面给带`@SaCheckLogin`这样的接口发请求的时候，把token放header里传过来就行。

### 7.前端代码

为了方便你的理解和使用，我特意写了html代码供你测试使用，代码已经写了大篇幅的注释，我就不讲解了。

```html
<!DOCTYPE html>
<html>
<head>
	<title>My Page</title>
	<meta charset="UTF-8">
	<!-- 引用jQuery库 -->
	<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.5.1/jquery.min.js"></script>
	<script>
		$(document).ready(function() {
			// 当页面加载完成时执行以下代码
			$("#login").click(function() {
				// 当“Login”按钮被点击时执行以下代码
				var username = $("#username").val(); // 获取输入框中的用户名
				var password = $("#password").val(); // 获取输入框中的密码
				var url = "http://localhost:8081/api/user/login?username=" + username + "&password=" + password; // 构造用于登录的URL
				$.get(url, function(data) { // 发送GET请求
					if (data.code == 200) { // 如果返回码为200则表示登录成功
						localStorage.setItem("token", data.data); // 将token保存到localStorage中
						alert("Login successful!"); // 弹出登录成功提示框
					} else {
						alert("Login failed: " + data.msg); // 弹出登录失败提示框
					}
				});
			});

			$("#get-data").click(function() {
				// 当“Get Data”按钮被点击时执行以下代码
				var token = localStorage.getItem("token"); // 获取localStorage中保存的token
				if (token == null || token == "") {
					alert("Please login first!"); // 如果token为空则表示未登录，弹出提示框
					return;
				}
				var url = "http://localhost:8081/api/user"; // 构造请求数据的URL
				$.ajax({
					url: url,
					type: "GET",
					headers: {
						"token": token // 在请求头中添加token
					},
					success: function(data) { // 请求成功回调函数
						$("#result").val(JSON.stringify(data)); // 将返回的数据展示到文本框中
					},
					error: function(xhr, status, error) { // 请求失败回调函数
						alert("Error: " + error); // 弹出错误提示框
					}
				});
			});
		});
	</script>
</head>
<body>
	<h1>Welcome to My Page</h1>
	<form>
		<label for="username">Username:</label>
		<input type="text" name="username" id="username"><br><br>
		<label for="password">Password:</label>
		<input type="password" name="password" id="password"><br><br>
		<input type="button" value="Login" id="login"> <!-- 点击该按钮发起登录请求 -->
	</form>
	<br>
	<textarea id="result" rows="10" cols="50"></textarea><br>
	<input type="button" value="Get Data" id="get-data"> <!-- 点击该按钮发起获取数据请求 -->
</body>
</html>
```
## 结语

以上就是精华内容，如果你不太理解，还是建议你，把GitHub上的代码[🔗https://github.com/PerCheung/learnsatoken](https://github.com/PerCheung/learnsatoken)拉下来实操一遍。光看不如上手做一遍。对了，我的user表sql如下。如果你想要了解shiro五表的使用，方便你理解sa-token对权限的管理，我也有shiro的开源项目链接[🔗https://github.com/PerCheung/learnshiro](https://github.com/PerCheung/learnshiro)。

```sql
drop database shiro_learn;

create database shiro_learn;

use shiro_learn;

create table user
(
    username    varchar(36) primary key comment '用户名',
    password    varchar(32) not null comment '密码',
    salt        varchar(36) not null comment '盐',
    create_time datetime default now() comment '创建时间',
    update_time datetime default now() comment '修改时间',
    deleted     int      default 0 comment '逻辑删除'
) comment '用户表'
    engine = innodb
    default charset = utf8mb4;

create trigger user_update
    before update
    on user
    for each row set new.update_time = now();
```
