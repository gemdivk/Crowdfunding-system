
document.addEventListener("DOMContentLoaded", () => {
    const donationForm = document.getElementById("donation-form");

    if (donationForm) {
        donationForm.addEventListener("submit", async (event) => {
            event.preventDefault();
            const urlParams = new URLSearchParams(window.location.search);
            const campaignId = urlParams.get("id");
            console.log(campaignId)
            const amount = parseFloat(document.getElementById("amount").value);
            if (!campaignId) {
                console.error("Campaign ID is missing in the URL");
            }

            const token = localStorage.getItem("token");
            if (!token) {
                alert("You must be logged in to donate.");
                window.location.href = "login.html";
                return;
             }
            const decoded = jwt_decode(token);
            const userID = decoded.UserID;
            console.log("user id is:", userID);
            const response = await fetch(`http://localhost:8080/campaigns/${campaignId}/donations/`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": `Bearer ${token}`
                },
                body: JSON.stringify({ amount, user_id: userID, })
            });
            console.log(response); // Log the full response
            if (response.ok) {
                alert("Donation successful!");
                window.location.href = "/static/campaign.html";
            } else {
                alert("Donation failed");
            }
        });
    }
});
