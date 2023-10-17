import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.*;

import java.util.ArrayList;
import java.util.List;

@SpringBootApplication
@RestController
public class MessageApi {

    private List<String> messages = new ArrayList<>();

    @GetMapping("/messages")
    public List<String> getMessages() {
        return messages;
    }

    @PostMapping("/messages")
    public void addMessage(@RequestBody String message) {
        messages.add(message);
    }

    public static void main(String[] args) {
        SpringApplication.run(MessageApi.class, args);
    }
}
