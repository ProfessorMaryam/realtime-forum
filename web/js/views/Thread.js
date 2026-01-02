export default async function ThreadView() {
    const params = new URLSearchParams(location.search);
    const postId = params.get("post");

    if (!postId) {
        return `<div class="home"><p>Missing post id.</p></div>`;
    }

    let comments = [];

    try {
        const res = await fetch(`/api/comments?post=${encodeURIComponent(postId)}`);
        if (res.ok) {
            comments = await res.json();
        }
    } catch {
        comments = [];
    }

    function escapeHtml(str) {
        if (!str) return "";
        return String(str).replace(/[&<>'\"]/g, (m) => ({
            '&': '&amp;',
            '<': '&lt;',
            '>': '&gt;',
            "'": '&#39;',
            '"': '&quot;'
        }[m]));
    }

    // Comments section content (this is the ONLY conditional part)
    const commentsHtml = comments.length > 0
        ? comments.map(c => `
            <div class="comment-block">
                <div class="comment-user">
                    <div class="username">
                        <h2>${escapeHtml(c.author)}</h2>
                    </div>
                </div>
                <div class="content1">
                    <pre>${escapeHtml(c.content)}</pre>
                </div>
            </div>
        `).join("")
        : `<p class="no-comments">No comments yet. Be the first to comment!</p>`;

    return `
        <!-- POST (always visible) -->
        <div class="focus-post">
            <h2>@author</h2>

            <div class="frame">
                <h3>Post title</h3>
                <div class="content">
                    <pre>Post content placeholder</pre>
                </div>
            </div>
        </div>

        <!-- COMMENT FORM (always visible) -->
        <div class="comment-input-container">
            <div class="comment-user">
                <div class="username">
                    <h2>@you</h2>
                </div>
            </div>

            <form class="comment-input-section">
                <input type="hidden" name="post_id" value="${escapeHtml(postId)}">
                <textarea
                    name="content"
                    placeholder="Write your comment..."
                    required
                ></textarea>
                <button class="post-comment-btn">Post Comment</button>
            </form>
        </div>

        <!-- COMMENT COUNT (always visible) -->
        <div class="comment-count">
            <span>${comments.length} Comments</span>
        </div>

        <!-- COMMENTS LIST (content changes) -->
        <div id="comments-list">
            ${commentsHtml}
        </div>
    `;
}
