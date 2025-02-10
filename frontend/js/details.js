const urlParams = new URLSearchParams(window.location.search);
const campaignId = urlParams.get("id");

async function loadCampaignDetails() {
    try {
        const response = await fetch(`/campaigns/${campaignId}`); // Adjust API URL
        const campaign = await response.json();

        document.getElementById("campaign-title").innerText = campaign.title;
        document.getElementById("campaign-description").innerText = campaign.description;
        document.getElementById("campaign-goal").innerText = campaign.target_amount;
        document.getElementById("campaign-raised").innerText = campaign.amount_raised;

        const mediaContainer = document.getElementById("campaign-media");
        mediaContainer.innerHTML = `<img src="${campaign.media_path}" alt="Campaign Image" class="campaign-img">`;
        const goal = Number(campaign.target_amount);
        const raised = Number(campaign.amount_raised);
        const progressPercentage = Math.min((raised / goal) * 100, 100);
        document.getElementById("progress-bar").style.width = progressPercentage + "%";
        document.getElementById("progress-percent").innerText = progressPercentage.toFixed(1) + "%";


        document.getElementById("donate-btn").addEventListener("click", () => {
            window.location.href = `/static/donation.html?id=${campaignId}`;
        });

    } catch (error) {
        console.error("Error loading campaign:", error);
    }
}

loadCampaignDetails();
