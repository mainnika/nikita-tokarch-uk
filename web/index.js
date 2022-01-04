import "code-prettify/styles/desert.css";
import "bootstrap-icons/font/bootstrap-icons.css";

import "code-prettify";

import "./style/index.scss";

const applyPrettyPrint = () => {

    const blogPosts = document.getElementsByClassName("blog-post-content");
    for (const blogPost of blogPosts) {

        const codeBlocks = blogPost.getElementsByTagName("code");
        for (const codeBlock of codeBlocks) {
            codeBlock.parentElement.classList.add("prettyprint", "linenums", "kg-card", "kg-code-card");
        }
    }

    PR.prettyPrint();
}

const init = () => {
    applyPrettyPrint();    
}

if (document.readyState === "complete" || document.readyState === "interactive") {
    requestAnimationFrame(init);
} else {
    document.addEventListener("DOMContentLoaded", init);
}




