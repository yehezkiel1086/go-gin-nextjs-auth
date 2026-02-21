"use client";

import { useState, useEffect } from "react";
import { JobCard } from "@/components/JobCard";
import { Button } from "@/components/ui/button";
import { Alert, AlertDescription } from "@/components/ui/alert";
import { getJobs, type Job } from "@/lib/api";
import { Briefcase, Plus } from "lucide-react";
import Link from "next/link";

const JobsPage = () => {
  const [jobs, setJobs] = useState<Job[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    fetchJobs();
  }, []);

  const fetchJobs = async () => {
    try {
      setIsLoading(true);
      setError("");
      const response = await getJobs();
      
      if (response.jobs) {
        setJobs(response.jobs);
      }
    } catch (err) {
      setError(err instanceof Error ? err.message : "Failed to fetch jobs. Please try again.");
    } finally {
      setIsLoading(false);
    }
  };

  if (isLoading) {
    return (
      <div className="container mx-auto py-8">
        <div className="flex items-center justify-center min-h-100">
          <div className="text-center">
            <Briefcase className="w-12 h-12 mx-auto mb-4 text-gray-400 animate-pulse" />
            <p className="text-gray-600">Loading jobs...</p>
          </div>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="container mx-auto py-8">
        <div className="max-w-2xl mx-auto">
          <Alert variant="destructive">
            <AlertDescription>{error}</AlertDescription>
          </Alert>
          <div className="mt-4 text-center">
            <Button onClick={fetchJobs} variant="outline">
              Try Again
            </Button>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="container mx-auto py-8">
      <div className="flex justify-between items-center mb-8">
        <div>
          <h1 className="text-3xl font-bold text-gray-900 mb-2">
            Available Jobs
          </h1>
          <p className="text-gray-600">
            {jobs.length} {jobs.length === 1 ? 'job' : 'jobs'} available
          </p>
        </div>
        <Link href="/jobs/create" className="flex items-center gap-2 bg-black text-white px-4 py-2 rounded-md hover:bg-gray-600 transition-colors">
          <Plus className="w-4 h-4" />
          Post a Job
        </Link>
      </div>

      {jobs.length === 0 ? (
        <div className="text-center py-16">
          <Briefcase className="w-16 h-16 mx-auto mb-4 text-gray-300" />
          <h2 className="text-xl font-semibold text-gray-700 mb-2">
            No jobs available
          </h2>
          <p className="text-gray-500 mb-6">
            Be the first to post a job opportunity!
          </p>
          <Button className="flex items-center gap-2 mx-auto">
            <Plus className="w-4 h-4" />
            Post a Job
          </Button>
        </div>
      ) : (
        <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
          {jobs.map((job) => (
            <JobCard key={job.ID} job={job} />
          ))}
        </div>
      )}
    </div>
  );
};

export default JobsPage;
