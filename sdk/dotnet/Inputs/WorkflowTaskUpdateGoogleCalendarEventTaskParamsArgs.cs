// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rootly.Inputs
{

    public sealed class WorkflowTaskUpdateGoogleCalendarEventTaskParamsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Days to adjust meeting by
        /// </summary>
        [Input("adjustmentDays")]
        public Input<int>? AdjustmentDays { get; set; }

        [Input("attendees")]
        private InputList<string>? _attendees;

        /// <summary>
        /// Emails of attendees
        /// </summary>
        public InputList<string> Attendees
        {
            get => _attendees ?? (_attendees = new InputList<string>());
            set => _attendees = value;
        }

        /// <summary>
        /// Value must be one of true or false
        /// </summary>
        [Input("canGuestsInviteOthers")]
        public Input<bool>? CanGuestsInviteOthers { get; set; }

        /// <summary>
        /// Value must be one of true or false
        /// </summary>
        [Input("canGuestsModifyEvent")]
        public Input<bool>? CanGuestsModifyEvent { get; set; }

        /// <summary>
        /// Value must be one of true or false
        /// </summary>
        [Input("canGuestsSeeOtherGuests")]
        public Input<bool>? CanGuestsSeeOtherGuests { get; set; }

        /// <summary>
        /// Sets the video conference type attached to the meeting. Value must be one of `eventHangout`, `eventNamedHangout`, `hangoutsMeet`, `addOn`.
        /// </summary>
        [Input("conferenceSolutionKey")]
        public Input<string>? ConferenceSolutionKey { get; set; }

        /// <summary>
        /// The event description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The event ID
        /// </summary>
        [Input("eventId", required: true)]
        public Input<string> EventId { get; set; } = null!;

        /// <summary>
        /// Meeting duration in format like '1 hour', '30 minutes'
        /// </summary>
        [Input("meetingDuration")]
        public Input<string>? MeetingDuration { get; set; }

        /// <summary>
        /// Value must be one of true or false
        /// </summary>
        [Input("postToIncidentTimeline")]
        public Input<bool>? PostToIncidentTimeline { get; set; }

        [Input("postToSlackChannels")]
        private InputList<Inputs.WorkflowTaskUpdateGoogleCalendarEventTaskParamsPostToSlackChannelArgs>? _postToSlackChannels;
        public InputList<Inputs.WorkflowTaskUpdateGoogleCalendarEventTaskParamsPostToSlackChannelArgs> PostToSlackChannels
        {
            get => _postToSlackChannels ?? (_postToSlackChannels = new InputList<Inputs.WorkflowTaskUpdateGoogleCalendarEventTaskParamsPostToSlackChannelArgs>());
            set => _postToSlackChannels = value;
        }

        /// <summary>
        /// Value must be one of true or false
        /// </summary>
        [Input("replaceAttendees")]
        public Input<bool>? ReplaceAttendees { get; set; }

        /// <summary>
        /// Send an email to the attendees notifying them of the event. Value must be one of true or false
        /// </summary>
        [Input("sendUpdates")]
        public Input<bool>? SendUpdates { get; set; }

        /// <summary>
        /// The event summary
        /// </summary>
        [Input("summary")]
        public Input<string>? Summary { get; set; }

        [Input("taskType")]
        public Input<string>? TaskType { get; set; }

        /// <summary>
        /// Time of meeting in format HH:MM
        /// </summary>
        [Input("timeOfMeeting")]
        public Input<string>? TimeOfMeeting { get; set; }

        public WorkflowTaskUpdateGoogleCalendarEventTaskParamsArgs()
        {
        }
        public static new WorkflowTaskUpdateGoogleCalendarEventTaskParamsArgs Empty => new WorkflowTaskUpdateGoogleCalendarEventTaskParamsArgs();
    }
}
