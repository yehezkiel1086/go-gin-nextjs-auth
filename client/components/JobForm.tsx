"use client";

import { useState } from "react";
import { faker } from '@faker-js/faker';
import { Button } from "@/components/ui/button"
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Textarea } from "@/components/ui/textarea"
import { Alert, AlertDescription } from "@/components/ui/alert"
import { useRouter } from "next/navigation";

interface CreateJobRequest {
  title: string;
  description: string;
  company: string;
  location: string;
  salary: number;
}

export function JobForm() {
  const [formData, setFormData] = useState<CreateJobRequest>({
    title: "",
    description: "",
    company: "",
    location: "",
    salary: 0,
  });
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState("");
  const [success, setSuccess] = useState(false);
  const router = useRouter()

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value, type } = e.target;
    setFormData(prev => ({
      ...prev,
      [name]: type === 'number' ? (value === '' ? 0 : Number(value)) : value,
    }));
    if (error) setError("");
  };

  const handleAutofill = () => {
    const jobTitles = [
      'Software Engineer',
      'Machine Learning Engineer',
      'Product Manager',
      'UX Designer',
      'Data Scientist',
      'DevOps Engineer',
      'Frontend Developer',
      'Backend Developer',
      'Full Stack Developer',
      'Mobile App Developer'
    ];
    
    const companies = [
      'GDP Labs',
      'Tokopedia',
      'Gojek',
      'Traveloka',
      'Bukalapak',
      'Shopee',
      'Lazada',
      'Grab',
      'OVO',
      'DANA'
    ];
    
    const locations = [
      'Jakarta',
      'Surabaya',
      'Bandung',
      'Yogyakarta',
      'Medan',
      'Semarang',
      'Makassar',
      'Palembang',
      'Tangerang',
      'Depok'
    ];

    setFormData({
      title: faker.helpers.arrayElement(jobTitles),
      description: faker.lorem.words(20),
      company: faker.helpers.arrayElement(companies),
      location: faker.helpers.arrayElement(locations),
      salary: faker.number.int({ min: 5000000, max: 25000000 }),
    });
    setError("");
    setSuccess(false);
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    
    if (!formData.title || !formData.description || !formData.company || !formData.location) {
      setError("Please fill in all fields");
      return;
    }

    if (formData.salary <= 0) {
      setError("Please enter a valid salary amount");
      return;
    }

    setIsLoading(true);
    setError("");

    try {
      const response = await fetch(`${process.env.NEXT_PUBLIC_API_URI}/jobs`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        credentials: "include",
        body: JSON.stringify({
          title: formData.title,
          description: formData.description,
          company: formData.company,
          location: formData.location,
          salary: formData.salary,
        })
      });

      if (!response.ok) {
        const errorData = await response.json().catch(() => ({}));
        throw new Error(errorData.message || `HTTP error! status: ${response.status}`);
      }

      router.push("/jobs");
    } catch (err) {
      setError(err instanceof Error ? err.message : "Failed to create job. Please try again.");
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <Card className="w-full max-w-2xl mx-auto">
      <CardHeader>
        <CardTitle>Create New Job Posting</CardTitle>
        <CardDescription>
          Fill in the details below to create a new job opportunity
        </CardDescription>
      </CardHeader>
      <CardContent>
        <form onSubmit={handleSubmit}>
          <div className="flex flex-col gap-6">
            <div className="grid gap-2">
              <Label htmlFor="title">Job Title</Label>
              <Input
                id="title"
                name="title"
                type="text"
                placeholder="e.g. Machine Learning Engineer"
                value={formData.title}
                onChange={handleInputChange}
                required
                disabled={isLoading}
              />
            </div>
            
            <div className="grid gap-2">
              <Label htmlFor="description">Description</Label>
              <Textarea
                id="description"
                name="description"
                placeholder="Describe the role, responsibilities, and requirements..."
                value={formData.description}
                onChange={handleInputChange}
                required
                disabled={isLoading}
                rows={4}
              />
            </div>

            <div className="grid gap-2">
              <Label htmlFor="company">Company</Label>
              <Input
                id="company"
                name="company"
                type="text"
                placeholder="e.g. GDP Labs"
                value={formData.company}
                onChange={handleInputChange}
                required
                disabled={isLoading}
              />
            </div>

            <div className="grid gap-2">
              <Label htmlFor="location">Location</Label>
              <Input
                id="location"
                name="location"
                type="text"
                placeholder="e.g. Surabaya"
                value={formData.location}
                onChange={handleInputChange}
                required
                disabled={isLoading}
              />
            </div>

            <div className="grid gap-2">
              <Label htmlFor="salary">Salary (IDR)</Label>
              <Input
                id="salary"
                name="salary"
                type="number"
                placeholder="e.g. 12000000"
                value={formData.salary || ""}
                onChange={handleInputChange}
                required
                disabled={isLoading}
                min="0"
              />
            </div>

            {error && (
              <Alert variant="destructive">
                <AlertDescription>
                  {error}
                </AlertDescription>
              </Alert>
            )}

            {success && (
              <Alert variant="success">
                <AlertDescription>
                  Job posting created successfully!
                </AlertDescription>
              </Alert>
            )}
          </div>
        </form>
      </CardContent>
      <CardFooter className="flex gap-2">
        <Button 
          type="submit" 
          onClick={handleSubmit}
          disabled={isLoading}
          className="flex-1"
        >
          {isLoading ? "Creating Job..." : "Create Job Posting"}
        </Button>
        <Button 
          type="button" 
          variant="outline"
          onClick={handleAutofill}
          disabled={isLoading}
        >
          Autofill
        </Button>
        <Button 
          type="button" 
          variant="outline"
          onClick={() => {
            setFormData({
              title: "",
              description: "",
              company: "",
              location: "",
              salary: 0,
            });
            setError("");
            setSuccess(false);
          }}
          disabled={isLoading}
        >
          Clear
        </Button>
      </CardFooter>
    </Card>
  )
}
